package resilience

import (
	"context"
	"time"
)

// WithRetry executes a function with exponential backoff
func WithRetry(ctx context.Context, attempts int, sleep time.Duration, f func() error) error {
	if err := f(); err != nil {
		if attempts--; attempts > 0 {
			// Exponential backoff
			jitter := sleep + (time.Duration(attempts) * time.Millisecond * 100)
			select {
			case <-time.After(jitter):
				return WithRetry(ctx, attempts, sleep*2, f)
			case <-ctx.Done():
				return ctx.Err()
			}
		}
		return err
	}
	return nil
}