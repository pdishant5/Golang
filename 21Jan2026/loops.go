package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	// conventional for loop
	for i := 1; i <= 10; i++ {
		fmt.Println(n, "*", i, "=", n*i)
	}

	// while-styles for loop
	i := 1
	for i <= 10 {
		fmt.Println(n, "*", i, "=", n*i)
		i++
	}

	// infinite for loop
	i = 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(n, "*", i, "=", n*i)
		i++
	}

	// for loop with range
	indeces := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, val := range indeces {
		fmt.Println(n, "*", i+1, "=", n*val)
	}
}
