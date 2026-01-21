package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	// no need of break statements after each case - implicit
	// switch without condition - all cases have to be boolean expressions
	switch {
	case age < 18:
		fmt.Println("You are minor!")
	case age >= 18:
		fmt.Println("You are adult!")
	}

	var choice int
	fmt.Print("Enter Your Choice: ")
	fmt.Scan(&choice)

	// explicit fallthrough to execute following case of the first matching case
	switch choice {
	case 1:
		fmt.Println("You chose 1!")
		fallthrough // case 2 will also be executed
		// fmt.Println("I am unreachable!")
	case 2:
		fmt.Println("You chose 2!")
	case 3, 4: // is executed if any value mathces
		fmt.Println("You chose 3!")
	default:
		fmt.Println("Invalid Choice!")
	}
}
