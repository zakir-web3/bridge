package scheduler

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRun_StopsAfterMaxErrors(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := RunWithConfig(ctx, &Config{
		BaseInterval:      time.Millisecond,
		MaxErrorCount:     3,
		MaxSleepTime:      5 * time.Millisecond,
		BackoffMultiplier: 1.1,
		RecoveryThreshold: 10,
	}, func(context.Context) error {
		return errors.New("always fail")
	})
	if err == nil {
		t.Fatal("expected scheduler to stop after max errors")
	}
	if want := "reached maximum error count 3"; err.Error() != want+", stopping scheduler" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRun_StopsOnCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()

	err := RunWithConfig(ctx, &Config{
		BaseInterval:      5 * time.Millisecond,
		MaxErrorCount:     100,
		MaxSleepTime:      time.Second,
		BackoffMultiplier: 1.1,
		RecoveryThreshold: 3,
	}, func(context.Context) error {
		return nil
	})
	if err != nil {
		t.Fatalf("expected nil on cancel, got %v", err)
	}
}
