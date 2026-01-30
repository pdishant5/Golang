package main

import "fmt"

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	go func() {
		channel1 <- "Hello from the goroutine 1!"
	}()
	go func() {
		channel2 <- "Hello from the goroutine 2!"
	}()
	go func() {
		channel3 <- "Hello from the goroutine 3!"
	}()

	select {
	case msg1 := <-channel1:
		fmt.Println("Received from channel1:", msg1)
	case msg2 := <-channel2:
		fmt.Println("Received from channel2:", msg2)
	case msg3 := <-channel3:
		fmt.Println("Received from channel3:", msg3)
	}

	// message1 := <-channel1 // receives message from the channel
	// message2 := <-channel2
	// message3 := <-channel3 // no fixed order of receiving messages - they are received in the order they are sent

	// fmt.Println(message1)
	// fmt.Println(message2)
	// fmt.Println(message3)
}
