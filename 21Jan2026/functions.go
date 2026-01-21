package main

import "fmt"

func add(a int, b int) int {
	sum := a + b
	return sum
}

// we can also write the above function as following
// func add(a, b int) (int) {
// 	sum := a + b
// 	return sum
// }

// named return values
// func add(a, b int) (sum int) {
// 	sum = a + b
// 	return sum // or return
// }

// variadic arguments (var-args)
func addNumbers(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func apply(f func(int) int, x int) int {
	return f(x)
}

func main() {
	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a)
	fmt.Scan(&b)

	sum := add(a, b)
	fmt.Println("The sum of", a, "and", b, "is:", sum)

	sum1 := addNumbers(1, 2, 3, 4, 5)
	sum2 := addNumbers(15, 19, 70)
	sum3 := addNumbers(10, 20)

	nums := []int{10, 20, 30, 40}
	sum4 := addNumbers(nums...) // passing slice to variadic functions

	fmt.Println("Sum1:", sum1)
	fmt.Println("Sum2:", sum2)
	fmt.Println("Sum3:", sum3)
	fmt.Println("Sum4:", sum4)
	fmt.Println("Sum4:", addNumbers())

	// functions as values
	ans := apply(func(x int) int {
		return x * 2
	}, 10)
	fmt.Println("Answer:", ans)

	// Anonymous functions - immediate invocations
	func() {
		fmt.Println("Hello, Dishant!")
	}()
}
