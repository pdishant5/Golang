package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for number := range jobs {
		fmt.Printf("Worker %d processing number %d!\n", id, number)
		time.Sleep(time.Millisecond * 500) // simulate work
		fmt.Printf("Worker %d finished number %d (square = %d).\n", id, number, number*number)
	}
}

func main() {
	jobs := make(chan int) // channel for sending jobs
	var wg sync.WaitGroup

	workerCount := 3

	// Start worker goroutines
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for i := 1; i <= 6; i++ {
		jobs <- i
	}

	close(jobs) // no more jobs
	wg.Wait()   // wait for all workers to finish

	fmt.Println("All jobs processed!")
}
