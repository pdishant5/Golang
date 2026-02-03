package main

import (
	"fmt"
	"sync"
)

func recurse(n int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Enter level", n)

	if n == 0 {
		fmt.Println("Base work at level 0")
		return
	}

	wg.Add(1)
	go recurse(n-1, wg) // recursive goroutine
	// wg.Wait() // deadlock occurs
	fmt.Println("Exit level", n)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go recurse(3, &wg)
	wg.Wait()

	fmt.Println("Main done")
}
