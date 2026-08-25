package scheduler

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/rs/zerolog/log"
)

const (
	MaxSleep            = 60 * time.Second
	MaxErrCount         = 10
	BaseBackoff         = 1.4
	MinRecoveryInterval = 3 // Start recovery after 3 consecutive successful executions
)

// Config provides configuration options for the scheduler
type Config struct {
	MaxErrorCount     int           // Maximum number of errors before stopping
	BaseInterval      time.Duration // Base interval between task executions
	MaxSleepTime      time.Duration // Maximum sleep time during backoff
	BackoffMultiplier float64       // Multiplier for exponential backoff
	RecoveryThreshold int           // Threshold for recovery after errors
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	return &Config{
		MaxErrorCount:     MaxErrCount,
		BaseInterval:      1 * time.Second,
		MaxSleepTime:      MaxSleep,
		BackoffMultiplier: BaseBackoff,
		RecoveryThreshold: MinRecoveryInterval,
	}
}

func Run(ctx context.Context, interval time.Duration, taskFunc func(ctx context.Context) error) error {
	return RunWithConfig(ctx, &Config{
		BaseInterval:      interval,
		MaxErrorCount:     MaxErrCount,
		MaxSleepTime:      MaxSleep,
		BackoffMultiplier: BaseBackoff,
		RecoveryThreshold: MinRecoveryInterval,
	}, taskFunc)
}

func RunWithConfig(ctx context.Context, config *Config, taskFunc func(ctx context.Context) error) error {
	errCount := 0
	successCount := 0
	sleepTime := config.BaseInterval

	timer := time.NewTimer(sleepTime)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			start := time.Now()
			err := taskFunc(ctx)
			executionTime := time.Since(start)

			if err != nil {
				errCount++
				successCount = 0 // Reset success count

				// Exponential backoff strategy
				sleepTime = calculateBackoff(config.BaseInterval, errCount, config.BackoffMultiplier, config.MaxSleepTime)

				log.Error().Err(err).
					Int("attempt", errCount).
					Dur("execution_time", executionTime).
					Dur("next_delay", sleepTime).
					Msg("Task execution failed")

				// Reached maximum error count, return error
				if errCount >= config.MaxErrorCount {
					log.Error().
						Int("max_error_count", config.MaxErrorCount).
						Msg("Reached maximum error count, stopping scheduler")
					return fmt.Errorf("reached maximum error count %d, stopping scheduler", config.MaxErrorCount)
				}
			} else {
				successCount++
				log.Debug().
					Dur("execution_time", executionTime).
					Int("success_count", successCount).
					Msg("Task executed successfully")

				// Only start recovery after reaching the threshold of consecutive successes
				if errCount > 0 && successCount >= config.RecoveryThreshold {
					errCount = 0
					successCount = 0
					sleepTime = config.BaseInterval
					log.Info().
						Dur("interval", sleepTime).
						Msg("Task executed successfully multiple times, restored normal interval")
				} else if errCount > 0 {
					// Reduce error count but maintain backoff
					errCount--
					sleepTime = calculateBackoff(config.BaseInterval, errCount, config.BackoffMultiplier, config.MaxSleepTime)
					log.Info().
						Int("error_count", errCount).
						Int("success_count", successCount).
						Dur("next_delay", sleepTime).
						Msg("Task executed successfully, error count reduced")
				} else {
					sleepTime = config.BaseInterval
				}
			}

			// Ensure next execution time doesn't execute immediately if task execution time exceeds sleep time
			if executionTime >= sleepTime {
				sleepTime = config.BaseInterval
			}

			// Reset timer
			timer.Reset(sleepTime)

		case <-ctx.Done():
			log.Info().Err(ctx.Err()).Msg("Scheduler received stop signal")
			return nil
		}
	}
}

// calculateBackoff calculates exponential backoff duration
func calculateBackoff(baseInterval time.Duration, errCount int, multiplier float64, maxSleep time.Duration) time.Duration {
	backoff := time.Duration(
		float64(baseInterval.Milliseconds())*math.Pow(multiplier, float64(errCount)),
	) * time.Millisecond

	if backoff > maxSleep {
		backoff = maxSleep
	}

	return backoff
}
