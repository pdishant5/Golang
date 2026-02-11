package main

import "fmt"

func main() {
	const x = 10

	var a int = x
	var b float64 = x
	// var c byte = x

	fmt.Printf("Type of x: %T;\n", x)
	fmt.Printf("a: %d; b: %0.2f\n", a, b)
	// fmt.Println(c)

	// var ab int = 10
	// f(ab)
}

// func f(x int64) {
// 	fmt.Println(x)
// }
