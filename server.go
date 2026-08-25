package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/zakir-web3/bridge/internal/bridge"
	"github.com/zakir-web3/bridge/internal/bridgehub"
	"github.com/zakir-web3/bridge/internal/cache"
	"github.com/zakir-web3/bridge/internal/scanner"
	"github.com/zakir-web3/bridge/internal/scheduler"
)

func Start(cfg *Config) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	errChan := make(chan error, 1)

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	log.Info().Str("version", Version).Msg("Starting bridge service...")
	log.Info().Str("bridge_interval", cfg.Bridge.Interval.String()).Msg("Bridge scanner configured")
	log.Info().Str("bridgeHub_interval", cfg.BridgeHub.Interval.String()).Msg("BridgeHub scanner configured")

	badgerCache, err := cache.NewBadgerCache(cfg.Source)
	if err != nil {
		return err
	}

	bridgeHubInstance, err := bridgehub.NewBridgeHub(ctx, cfg.BridgeHub)
	if err != nil {
		return err
	}
	bridgeInstance, err := bridge.NewBridge(ctx, cfg.Bridge, bridgeHubInstance)
	if err != nil {
		return err
	}
	bridgeHubInstance.SetBridgeContract(bridgeInstance)

	bridgeScanner := scanner.NewScanner(cfg.Bridge.Config, badgerCache, bridgeInstance)
	bridgeHubScanner := scanner.NewScanner(cfg.BridgeHub.Config, badgerCache, bridgeHubInstance)

	var wg sync.WaitGroup

	SafeGo("bridge-scanner", func() error {
		return scheduler.Run(ctx, cfg.Bridge.Interval, bridgeScanner.ScanBlockRange)
	}, errChan, &wg)

	SafeGo("bridgeHub-scanner", func() error {
		return scheduler.Run(ctx, cfg.BridgeHub.Interval, bridgeHubScanner.ScanBlockRange)
	}, errChan, &wg)

	SafeGo("finalize-withdrawals", func() error {
		if !cfg.Bridge.SendFinalizeWithdrawals || os.Getenv("ENABLE_FINALIZE_WITHDRAWALS") != "true" {
			log.Info().Msg("send_finalize_withdrawals is disabled, skipping finalize withdrawals")
			return nil
		}
		return scheduler.Run(ctx, cfg.Bridge.FinalizeWithdrawalsInterval, bridgeInstance.FinalizeWithdrawals)
	}, errChan, &wg)

	log.Info().Msg("All services started successfully")

	select {
	case err = <-errChan:
		cancelFunc()
		wg.Wait()
		return err
	case <-c:
		log.Info().Msg("received interrupt signal")
		cancelFunc()
		wg.Wait()
	}

	log.Info().Msg("Shutting down gracefully...")
	return nil
}

func SafeGo(name string, fn func() error, errChan chan<- error, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Error().
					Interface("panic", r).
					Str("goroutine", name).
					Str("stack", string(debug.Stack())).
					Msg("Goroutine panic recovered")

				// Send the panic as an error to the error channel
				errChan <- fmt.Errorf("goroutine %s panic: %v", name, r)
			}
			wg.Done()
		}()

		if err := fn(); err != nil {
			log.Error().Err(err).Str("goroutine", name).Msg("Goroutine error")
			errChan <- err
		}
	}()
}
