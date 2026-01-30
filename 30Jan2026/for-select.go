package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Learning concurrency in Go!")
		}
	}
}

func main() {
	charChannel := make(chan string, 3)
	chars := []string{"A", "B", "C"}

	for _, char := range chars {
		select {
		case charChannel <- char:
		}
	}

	close(charChannel)

	for char := range charChannel {
		println(char)
	}

	// go func() {
	// 	for {
	// 		select {
	// 		default:
	// 			fmt.Println("Learning concurrency in Go!")
	// 		}
	// 	}
	// }() // infinite running goroutine - terminated when parent goroutine terminates

	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Microsecond * 10)

	close(done)
}
