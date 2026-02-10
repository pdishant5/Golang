package main

import "fmt"

func main() {
	x := 1

	defer fmt.Println(x)
	defer func(y int) {
		fmt.Println(y)
	}(x)

	defer func(y *int) {
		fmt.Println(*y)
	}(&x)

	x++

	go Abc()
}
