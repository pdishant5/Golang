package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// close(doneChan)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(time.Second) // waits for a second to complete
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you're liking Go lang!", done)

	// <-done
	// <-done
	// <-done
	// <-done

	// time.Sleep(3 * time.Second) // waits to complete all goroutines
	// for _, done := range dones {
	// 	<-done
	// }

	// for doneChan := range done {
	// 	fmt.Println(doneChan)
	// }
	for range done {
	}

	// pipeline example
	// input
	nums := []int{2, 3, 4, 7, 1}
	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	finalChannel := square(dataChannel)
	// stage 3
	for n := range finalChannel {
		fmt.Println(n)
	}
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- (n * n)
		}
		close(out)
	}()
	return out
}
