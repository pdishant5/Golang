package main

import (
	"fmt"
	"time"
)

func greet(name string) {
	fmt.Println("Hello,", name+"!")

	// i := 0
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}() // these goroutines run even if the parent goroutine exits
	}
	fmt.Println("Thank you! Visit again!")
}

func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}() // this deferred function call will be responsible for recovering from panic

	// defer func() {
	// 	recoverFunction()
	// }() // this deferred function call doesn't recover from panic

	// defer recover() // this also doesn't recover

	if b == 0 {
		panic("division by zero")
	}

	result := a / b
	fmt.Println("Division Result:", result)
}

func recoverFunction() {
	if r := recover(); r != nil {
		fmt.Println("Recovered again from panic:", r)
	}
}

func main() {

	fmt.Println("Program starts here!")
	go greet("Dishant")

	time.Sleep(time.Second)

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic:", r)
	// 	}
	// }()

	safeDivide(4, 0)
	// go safeDivide(4, 0) // this goroutine is not recovered by the defer in the creator goroutine

	// time.Sleep(time.Second)
}
