package worker

import (
	"context"
	"testing"
	"time"
)

func TestWorkerCompletion(t *testing.T) {
	ctx, cancel := context.WithCancel(rootCtx)
	t.Cleanup(cancel)

	result := make(chan string)

	StartWorker(ctx, result)

	select {
	case res := <-result:
		if res != "Work completed!" {
			t.Fatalf("Expected: 'Work completed!'; Got: '%q'", res)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("Test timed out!")
	}
}

func TestWorkerCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(rootCtx)

	result := make(chan string)
	StartWorker(ctx, result)

	cancel()

	select {
	case res := <-result:
		if res != "Work cancelled!" {
			t.Fatalf("Expected: 'Work cancelled!'; Got: '%q'", res)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("Worker did not stop after context cancellation!")
	}
}
