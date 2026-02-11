package consumer

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"kafka-moduler/jobs"
	"kafka-moduler/results"
)

func Consume(ctx context.Context, id int, jobs <-chan jobs.Job, resultChan chan<- results.Result) {
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

			result := results.Result{
				JobID:  job.ID,
				Status: "processed",
			}
			resultChan <- result
		}
	}
}
