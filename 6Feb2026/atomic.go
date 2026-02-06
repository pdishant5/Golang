package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var activeRequests atomic.Int64

func handleRequest(wg *sync.WaitGroup) {
	defer wg.Done()

	// Request started
	activeRequests.Add(1)

	// Simulate work
	time.Sleep(100 * time.Millisecond)

	// Request finished
	activeRequests.Add(-1)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go handleRequest(&wg)
	}

	// Monitor active requests
	go func() {
		for {
			fmt.Println("Active requests:", activeRequests.Load())
			time.Sleep(50 * time.Millisecond)
		}
	}()

	wg.Wait()
	fmt.Println("All requests completed!")
}
