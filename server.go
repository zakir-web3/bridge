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
	"github.com/zakir-web3/bridge/internal/solana"
)

func Start(cfg *Config) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	errChan := make(chan error, 1)

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	log.Info().Str("version", Version).Msg("Starting bridge service...")
	if cfg.Solana.Enabled() {
		log.Info().Str("solana_interval", cfg.Solana.Interval.String()).Msg("Solana bridge scanner configured")
	}
	if cfg.BridgeHub.Enabled() {
		log.Info().Str("bridgeHub_interval", cfg.BridgeHub.Interval.String()).Msg("BridgeHub scanner configured")
	}
	if cfg.Bridge.Enabled() {
		log.Info().Str("bridge_interval", cfg.Bridge.Interval.String()).Msg("Bridge scanner configured")
	}

	badgerCache, err := cache.NewBadgerCache(cfg.Source)
	if err != nil {
		return err
	}

	bridgeHubInstance, err := bridgehub.NewBridgeHub(ctx, cfg.BridgeHub)
	if err != nil {
		return err
	}

	var bridgeInstance *bridge.Bridge
	if cfg.Bridge.Enabled() {
		bridgeInstance, err = bridge.NewBridge(ctx, cfg.Bridge, bridgeHubInstance)
		if err != nil {
			return err
		}
		bridgeHubInstance.SetBridgeContract(bridgeInstance)
	}

	var wg sync.WaitGroup

	if cfg.Solana.Enabled() {
		solanaBridgeInstance, err := solana.NewBridge(ctx, cfg.Solana, bridgeHubInstance)
		if err != nil {
			return err
		}
		solanaScanner := solana.NewScanner(solana.SlotScannerConfig{
			Interval:     cfg.Solana.Interval,
			StartSlot:    cfg.Solana.StartSlot,
			SlotInterval: cfg.Solana.SlotInterval,
			SlotDelay:    cfg.Solana.SlotDelay,
			ClearCache:   cfg.Solana.ClearCache,
		}, badgerCache, solanaBridgeInstance)
		SafeGo("solana-scanner", func() error {
			return scheduler.Run(ctx, cfg.Solana.Interval, solanaScanner.ScanSlotRange)
		}, errChan, &wg)
	}

	if cfg.Bridge.Enabled() {
		bridgeScanner := scanner.NewScanner(cfg.Bridge.Config, badgerCache, bridgeInstance)
		SafeGo("bridge-scanner", func() error {
			return scheduler.Run(ctx, cfg.Bridge.Interval, bridgeScanner.ScanBlockRange)
		}, errChan, &wg)

		SafeGo("finalize-withdrawals", func() error {
			if !cfg.Bridge.SendFinalizeWithdrawals || os.Getenv("ENABLE_FINALIZE_WITHDRAWALS") != "true" {
				log.Info().Msg("send_finalize_withdrawals is disabled, skipping finalize withdrawals")
				return nil
			}
			return scheduler.Run(ctx, cfg.Bridge.FinalizeWithdrawalsInterval, bridgeInstance.FinalizeWithdrawals)
		}, errChan, &wg)
	}

	bridgeHubScanner := scanner.NewScanner(cfg.BridgeHub.Config, badgerCache, bridgeHubInstance)
	SafeGo("bridgeHub-scanner", func() error {
		return scheduler.Run(ctx, cfg.BridgeHub.Interval, bridgeHubScanner.ScanBlockRange)
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
