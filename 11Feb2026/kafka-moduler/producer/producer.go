package producer

import (
	"context"
	"fmt"

	"kafka-moduler/jobs"
)

func Produce(ctx context.Context, jobChan chan<- jobs.Job) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping job producer!")
			close(jobChan)
			return
		case jobChan <- jobs.Job{ID: i, Data: fmt.Sprintf("Job %d!", i)}:
			fmt.Printf("Submitted job %d!\n", i)
			i++
			// time.Sleep(200 * time.Millisecond)
		}
	}
}
