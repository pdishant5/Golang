package processor

import (
	"context"
	"fmt"

	"kafka-moduler/results"
	"kafka-moduler/storage"
)

func Process(ctx context.Context, resultChan <-chan results.Result, store *storage.Storage) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping processor! Storing the already consumed jobs!")
			for result := range resultChan {
				store.Store(result.JobID, result.Status)
				fmt.Printf("Processor stored result for job %d!\n", result.JobID)
			}
			return
		case result, ok := <-resultChan:
			if !ok {
				fmt.Println("Processor result channel closed!")
				return
			}
			store.Store(result.JobID, result.Status)
			fmt.Printf("Processor stored result for job %d!\n", result.JobID)
		}
	}
}
