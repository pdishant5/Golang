package main

import (
	"context"
	"fmt"
	"time"
)

func fetchUserData(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("User data fetched successfully!")
		return nil

	case <-ctx.Done():
		// Context was cancelled or timed out
		return ctx.Err()
	}
}

func main() {
	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := fetchUserData(ctx)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}

	fmt.Println("Request completed!")
}
