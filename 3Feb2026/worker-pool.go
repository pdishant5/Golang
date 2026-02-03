package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Order represents a unit of work
type Order struct {
	ID     int
	Amount float64
}

// Result represents the result of processing
type Result struct {
	OrderID int
	Status  string
}

// Worker function
func worker(
	ctx context.Context,
	id int,
	jobs <-chan Order,
	results chan<- Result,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d shutting down!\n", id)
			return

		case order, ok := <-jobs:
			if !ok {
				// Channel closed → no more jobs
				fmt.Printf("Worker %d finished all jobs!\n", id)
				return
			}

			// Simulate processing time (I/O, API, DB, etc.)
			time.Sleep(time.Millisecond * time.Duration(200+rand.Intn(300)))

			fmt.Printf("Worker %d processed order %d!\n", id, order.ID)

			results <- Result{
				OrderID: order.ID,
				Status:  "processed",
			}
		}
	}
}

func main() {
	// rand.Seed(time.Now().UnixNano())

	// Root context (for graceful shutdown)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const workerCount = 4
	const totalOrders = 10

	jobs := make(chan Order, 5)     // buffered job channel
	results := make(chan Result, 5) // buffered results channel

	var wg sync.WaitGroup

	// Start worker pool
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, results, &wg)
	}

	// Producer: send orders
	go func() {
		for i := 1; i <= totalOrders; i++ {
			order := Order{
				ID:     i,
				Amount: rand.Float64() * 100,
			}
			jobs <- order
			fmt.Printf("Submitted order %d!\n", order.ID)
		}
		close(jobs) // important!
	}()

	// Close results channel when workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result received for order %d: %s!\n", result.OrderID, result.Status)
	}

	fmt.Println("All orders processed. Shutting down.")
}
