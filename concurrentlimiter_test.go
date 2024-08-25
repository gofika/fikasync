package fikasync

import (
	"testing"
	"time"
)

func TestConcurrentLimiter(t *testing.T) {
	// Create a new ConcurrentLimiter with a maximum of 2 concurrent operations.
	cl := NewConcurrentLimiter(2)

	// Run 3 operations concurrently.
	cl.Go(func() {
		time.Sleep(100 * time.Millisecond)
	})
	cl.Go(func() {
		time.Sleep(100 * time.Millisecond)
	})
	cl.Go(func() {
		time.Sleep(100 * time.Millisecond)
	})

	// Wait for all operations to complete.
	cl.Wait()

	// Output:
	// All operations completed
}
