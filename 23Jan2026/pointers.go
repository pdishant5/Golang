package main

import "fmt"

func main() {
	var x int = 42
	var p *int = &x // p is a pointer to an int, initialized to the address of x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of p (address of x):", p)
	fmt.Println("Value pointed to by p:", *p)

	changeValue(p)

	fmt.Println("New value of x after change:", x)
}

func changeValue(ptr *int) {
	*ptr = 100 // mutating the caller value
}
