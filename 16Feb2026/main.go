package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("1")
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}
	fmt.Println("Converted number:", n)

	m, err := strconv.ParseInt("01", 10, 64)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}
	fmt.Println("Converted number:", m)

	// a := strconv.Itoa(int(m))
	a := strconv.FormatInt(m, 10)
	fmt.Println("String:", a)

	b, err := strconv.ParseBool("true")
	fmt.Println("Boolean:", b)

	c, _ := fmt.Printf("Boolean: %t\n", b)
	fmt.Println("Number of bytes written:", c)

	str := "Hello, Gopher!"
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Character: %c\n", str[0])
	fmt.Printf("Type: %T\n", str[0]) // because string is a read only slice of bytes internally
	// str[0] = 'h' // compiule-time error - strings are immutable in Go
}
