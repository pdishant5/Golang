package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"kafka-moduler/consumer"
	"kafka-moduler/jobs"
	"kafka-moduler/processor"
	"kafka-moduler/producer"
	"kafka-moduler/results"
	"kafka-moduler/storage"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signalCh
		fmt.Printf("\nReceived signal: %s. Initiating graceful shutdown!\n\n", sig)
		cancel()
	}()

	jobs := make(chan jobs.Job, 5)
	results := make(chan results.Result, 5)

	store := storage.New()
	const consumerCount = 3
	var consumerWg sync.WaitGroup

	for i := 1; i <= consumerCount; i++ {
		consumerWg.Go(func() {
			consumer.Consume(ctx, i, jobs, results)
		})
	}

	var processorWg sync.WaitGroup

	processorWg.Go(func() {
		processor.Process(ctx, results, store)
	})

	go producer.Produce(ctx, jobs)

	func() {
		consumerWg.Wait()
		fmt.Println("Results channel closed!")
		close(results)
	}()

	processorWg.Wait()
	// close(jobs)
	fmt.Println("All jobs processed. Shutting down.")
	// time.Sleep(time.Second)

	fmt.Printf("\nAll components shut down gracefully\n\n")

	fmt.Println("Final in-memory data store:")
	store.Range(func(key int, value string) bool {
		fmt.Printf("Job %d: %s\n", key, value)
		return true
	})
}
