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

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Status string
}

type Storage struct {
	mu   sync.Mutex
	data map[int]string
}

func (s *Storage) Store(jobID int, status string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[jobID] = status
}

func (s *Storage) Range(fn func(JobId int, status string) bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for k, v := range s.data {
		fn(k, v)
	}
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[int]string),
	}
}

func Consumer(ctx context.Context, id int, jobs <-chan Job, results chan<- Result) {
	// defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Consumer %d shutting down!\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				// channel closed -> no more jobs
				fmt.Printf("Consumer %d finished all jobs!\n", id)
				return
			}

			time.Sleep(time.Millisecond * time.Duration(200+rand.Intn(300)))

			fmt.Printf("Consumer %d consumed job %d!\n", id, job.ID)

			result := Result{
				JobID:  job.ID,
				Status: "processed",
			}
			results <- result
		}
	}
}

func Producer(ctx context.Context, jobs chan<- Job) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping job producer!")
			return
		case jobs <- Job{ID: i, Data: fmt.Sprintf("Job %d!", i)}:
			fmt.Printf("Submitted job %d!\n", i)
			i++
			// time.Sleep(200 * time.Millisecond)
		}
	}
}

func Processor(ctx context.Context, results <-chan Result, storage *Storage) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping processor!")
			return
		case result, ok := <-results:
			if !ok {
				fmt.Println("Processor result channel closed!")
				return
			}
			storage.Store(result.JobID, result.Status)
			fmt.Printf("Processor stored result for job %d!\n", result.JobID)
		}
	}
}

func main1() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-signalCh
		fmt.Printf("\nReceived signal: %s. Initiating graceful shutdown!\n\n", sig)
		cancel()
	}()

	jobs := make(chan Job, 5)
	results := make(chan Result, 5)

	storage := NewStorage()
	const consumerCount = 3
	var consumerWg sync.WaitGroup

	for i := 1; i <= consumerCount; i++ {
		consumerWg.Go(func() {
			Consumer(ctx, i, jobs, results)
		})
	}

	var processorWg sync.WaitGroup

	processorWg.Go(func() {
		Processor(ctx, results, storage)
	})

	go Producer(ctx, jobs)

	go func() {
		consumerWg.Wait()
		close(results)
	}()

	processorWg.Wait()
	close(jobs)
	fmt.Println("All jobs processed. Shutting down.")

	fmt.Printf("\nAll components shut down gracefully\n\n")

	fmt.Println("Final in-memory data store:")
	storage.Range(func(key int, value string) bool {
		fmt.Printf("Job %d: %s\n", key, value)
		return true
	})
}
