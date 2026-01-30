package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	fmt.Println("Welcome,", name+"!")
}

func main() {
	go greet("Dishant")
	go greet("Vivek")
	go greet("Bhaumil") // no fuxed order of execution for goroutines

	time.Sleep(time.Second) // wait for goroutines to finish

	fmt.Println("Hello World!")
}
