package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

type TaskStatus int

const (
	Pending TaskStatus = iota
	Running
	Completed
	Failed
)

type Task struct {
	ID        int
	Content   string
	Status    TaskStatus
	CreatedAt time.Time
}

type Queue struct {
	queue []Task
	mu    sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{
		queue: []Task{},
	}
}

func (q *Queue) AddTask(task Task) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queue = append(q.queue, task)
}

func (q *Queue) NextTask() (Task, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.queue) == 0 {
		return Task{}, false
	}

	task := q.queue[0]
	q.queue = q.queue[1:]

	return task, true
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		fmt.Printf("\n>>> Received signal: %v. Starting graceful shutdown!<<<\n", sig)
		cancel()
	}()

	taskChan := make(chan Task, 10)

	const consumerCount = 5
	heap := NewQueue()
	var consumerWg sync.WaitGroup

	for i := 1; i <= consumerCount; i++ {
		consumerWg.Go(func() {
			Consumer(ctx, i, taskChan, heap)
		})
	}

	const processorCount = 5
	var processorWg sync.WaitGroup
	for i := 1; i <= processorCount; i++ {
		processorWg.Go(func() {
			Processor(ctx, i, heap)
		})
	}

	go Producer(ctx, taskChan)

	consumerWg.Wait()

	for _, task := range heap.queue {
		fmt.Println(task.ID, task.Content, task.Status)
	}
}

func Producer(ctx context.Context, taskChan chan<- Task) {
	id := 1

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Producer shutting down!")
			close(taskChan)
			return
		case <-ticker.C:
			task := Task{
				ID:        id,
				Content:   "Task " + strconv.Itoa(id) + "!",
				Status:    Pending,
				CreatedAt: time.Now(),
			}
			taskChan <- task
			fmt.Printf("Producer sent task %d!\n", id)

			id++
		}
	}
}

func Consumer(ctx context.Context, id int, taskChan <-chan Task, heap *Queue) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Consumer %d is shutting down!\n", id)
			return
		case task, ok := <-taskChan:
			if !ok {
				fmt.Printf("Consumer %d consumed all tasks!\n", id)
				return
			}
			heap.AddTask(task)
			fmt.Printf("Task %d has been added to the queue!\n", task.ID)
		}
	}
}

func Processor(ctx context.Context, id int, heap *Queue) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Processor %d is shutting down!\n", id)
			return
		default:
			task, next := heap.NextTask()
			if !next {
				fmt.Println("All tasks completed!")
				return
			}
			task.Status = Completed
		}
	}
}
