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

func NewStorage() *Storage {
	return &Storage{
		data: make(map[int]string),
	}
}

func Processor(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup, storage *Storage) {
	// defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d shutting down!\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				// channel closed -> no more jobs
				fmt.Printf("Worker %d finished all jobs!\n", id)
				return
			}

			time.Sleep(time.Millisecond * time.Duration(200+rand.Intn(300)))

			fmt.Printf("Worker %d processed order %d!\n", id, job.ID)

			result := Result{
				JobID:  job.ID,
				Status: "processed",
			}

			storage.Store(result.JobID, result.Status)

			results <- result
		}
	}
}

func main() {
	storage := NewStorage()

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

	const workerCount = 3
	var wg sync.WaitGroup

	for i := 1; i <= workerCount; i++ {
		wg.Go(func() {
			Processor(ctx, i, jobs, results, &wg, storage)
		})
	}

	go func() {
		i := 1
		for {
			job := Job{
				ID:   i,
				Data: fmt.Sprintf("I am the job %d!", i),
			}
			jobs <- job
			i++
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Result received for order %d: %s!\n", result.JobID, result.Status)
	}
	fmt.Println("All jobs processed. Shutting down.")
}
