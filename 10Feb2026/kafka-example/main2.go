package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*

=====================
Domain Models
=====================

*/

// Job represents a unit of work (same fields as Order)
// type Job struct {
// 	ID     int
// 	Amount float64
// }

// // Result represents processed output
// type Result struct {
// 	JobID  int
// 	Status string
// }

/*

=====================
Producer (Kafka Producer)
=====================

*/

func producer(ctx context.Context, jobs chan<- Job) {
	id := 1

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[Producer] shutting down")
			return
		case <-ticker.C:
			job := Job{
				ID:   id,
				Data: fmt.Sprintf("%.2f", rand.Float64()*100),
			}

			id++
			jobs <- job
			fmt.Printf("[Producer] produced job %d\n", job.ID)
		}
	}
}

/*

=====================
Processor (Kafka Stream Processor)
Fan-out pattern
=====================

*/

func processor(
	ctx context.Context,
	id int,
	input <-chan Job,
	output chan<- Result,
	wg *sync.WaitGroup,
) {
	// defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[Processor %d] shutting down\n", id)
			return
		case job, ok := <-input:
			if !ok {
				fmt.Printf("[Processor %d] input channel closed\n", id)
				return
			}

			// Simulate processing
			time.Sleep(time.Millisecond * time.Duration(200+rand.Intn(300)))

			result := Result{
				JobID:  job.ID,
				Status: "processed",
			}

			fmt.Printf("[Processor %d] processed job %d\n", id, job.ID)
			output <- result
		}
	}
}

/*

=====================
Consumer (Kafka Consumer)
Fan-in pattern
=====================

*/

func consumer(
	ctx context.Context,
	results <-chan Result,
	store *sync.Map,
	wg *sync.WaitGroup,
) {
	// defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[Consumer] shutting down")
			return
		case result, ok := <-results:
			if !ok {
				fmt.Println("[Consumer] results channel closed")
				return
			}

			// Store in in-memory datastore
			store.Store(result.JobID, result.Status)
			fmt.Printf("[Consumer] stored result for job %d\n", result.JobID)
		}
	}
}

/*

=====================
Main (Orchestrator)
=====================

*/

func main() {
	// rand.Seed(time.Now().UnixNano())

	// Root context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	// Listen for OS signals (Ctrl+C, SIGTERM)
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signalCh
		fmt.Printf("\n>>> Received signal: %s. Initiating graceful shutdown <<<\n", sig)
		cancel()
	}()

	/*
	   Topics (Kafka-like)
	*/
	jobsTopic := make(chan Job, 100)
	resultsTopic := make(chan Result, 100)

	/*
	   In-memory datastore (efficient & concurrent)
	*/
	var store sync.Map

	/*
	   Processor pool (fan-out)
	*/

	const processorCount = 4
	var processorWG sync.WaitGroup

	for i := 1; i <= processorCount; i++ {
		processorWG.Go(func() {
			processor(ctx, i, jobsTopic, resultsTopic, &processorWG)
		})
	}

	/*
	   Consumer group (fan-in)
	*/
	var consumerWG sync.WaitGroup

	consumerWG.Go(func() {
		consumer(ctx, resultsTopic, &store, &consumerWG)
	})

	/*
	   Producer
	*/
	go producer(ctx, jobsTopic)

	/*
	   Graceful shutdown coordination
	*/
	// Stop processors → then close results topic

	go func() {
		processorWG.Wait()
		close(resultsTopic)
	}()

	// Stop consumer
	consumerWG.Wait()

	// Close jobs topic last (producer exits via context)
	close(jobsTopic)
	fmt.Printf("\nAll components shut down gracefully\n")

	/*
	   Inspect in-memory store
	*/
	fmt.Println("Final in-memory data store:")

	store.Range(func(key, value any) bool {
		fmt.Printf("Job %d → %s\n", key.(int), value.(string))
		return true
	})
}
