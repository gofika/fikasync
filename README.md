[![codecov](https://codecov.io/gh/gofika/fikasync/branch/main/graph/badge.svg)](https://codecov.io/gh/gofika/fikasync)
[![Build Status](https://github.com/gofika/fikasync/workflows/build/badge.svg)](https://github.com/gofika/fikasync)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gofika/fikasync)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofika/fikasync)](https://goreportcard.com/report/github.com/gofika/fikasync)
[![Licenses](https://img.shields.io/github/license/gofika/fikasync)](LICENSE)

# fikasync

Utility designed for concurrent handling.


## Basic Usage

### Installation

To get the package, execute:

```bash
go get github.com/gofika/fikasync
```

### Example

```go
package main

import (
	"log"
	"time"

	"github.com/gofika/fikasync"
)

// fakeTask simulates a task that takes some time to complete.
func fakeTask(i int) {
	// your task code here...
	log.Printf("Executing task %d\n", i)
	// simulate some work
	time.Sleep(time.Second)
}

func main() {
	// create a new limiter that allows 5 concurrent operations
	limiter := fikasync.NewConcurrentLimiter(5)
	defer limiter.Close()

	// run 20 tasks
	for i := 0; i < 20; i++ {
		// run the task in a goroutine, blocking if the maximum number of concurrent operations has been reached
		limiter.Go(func() {
			fakeTask(i)
		})
	}

	limiter.Wait()
	log.Println("All tasks completed")
}
```