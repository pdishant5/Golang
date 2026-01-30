package main

import "fmt"

func main() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "Hello from the goroutine 1!"
	}()
	go func() {
		myChannel <- "Hello from the goroutine 2!"
	}()
	go func() {
		myChannel <- "Hello from the goroutine 3!"
	}()

	message1 := <-myChannel // receives message from the channel
	message2 := <-myChannel
	message3 := <-myChannel // no fixed order of receiving messages - they are received in the order they are sent

	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)
}
