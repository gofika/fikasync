package fikasync

import (
	"sync"
)

// ConcurrentLimiter is a simple concurrent limiter that limits the number of
type ConcurrentLimiter struct {
	limit chan struct{}
	wg    sync.WaitGroup
}

// NewConcurrentLimiter creates a new ConcurrentLimiter with the given maximum number of concurrent operations.
// If maxConcurrent is 0, the limiter is effectively disabled.
// Caller should call Close() when done with the limiter.
func NewConcurrentLimiter(maxConcurrent int) *ConcurrentLimiter {
	return &ConcurrentLimiter{
		limit: make(chan struct{}, maxConcurrent),
	}
}

// Go runs the given function in a goroutine, blocking if the maximum number of concurrent operations has been reached.
func (cl *ConcurrentLimiter) Go(fn func()) {
	cl.limit <- struct{}{}
	cl.wg.Add(1)
	go func() {
		defer func() {
			<-cl.limit
			cl.wg.Done()
		}()
		fn()
	}()
}

// Wait waits for all operations to complete.
func (cl *ConcurrentLimiter) Wait() {
	cl.wg.Wait()
}

// Close closes the limiter, causing all queued operations to be dropped.
// Caller should call this when done with the limiter.
func (cl *ConcurrentLimiter) Close() {
	close(cl.limit)
}
