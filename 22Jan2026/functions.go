package main

import "fmt"

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// fuction vs method example
type User struct {
	Name string
}

func (u User) Greet() string {
	return "Hello, " + u.Name + "!"
}

// pointer vs value receiver example
func (u *User) Rename(newName string) {
	u.Name = newName
}

// recover example with defer
func safeDivide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}() // this deferred function will be called but the panic will already be recovered
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered again from panic:", r)
		}
	}() // this deferred function will be recovering the panic as it is on the top in the stack

	if b == 0 {
		panic("division by zero")
	}

	result := a / b
	fmt.Println("Division Result:", result)
}

func main() {

	// closure example
	f := counter()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 2
	fmt.Println(f()) // 3

	g := counter()
	fmt.Println(g()) // 1
	fmt.Println(g()) // 2

	j := 0
	for j = 0; j < 3; j++ {
		defer func(j int) {
			fmt.Println("Deferred Call in Loop:", j)
		}(j)
	}
	i := 0
	for i = 0; i < 3; i++ {
		defer func() {
			fmt.Println("Deferred Call in Loop:", i)
		}()
	}
	for k := 0; k < 3; k++ {
		defer func() {
			fmt.Println("Deferred Call in Loop:", k)
		}()
	}

	// deferring function example
	defer fmt.Println("Deferred Call 1")
	defer fmt.Println("Deferred Call 2")
	fmt.Println("Main Function Execution 1")

	// arguments evaluated immediately
	x := 10
	defer fmt.Println("Deferred Call 3:", x) // 10 - evaluated immediately
	x = 20
	fmt.Println("Main Function Execution 2")

	// method example
	user := User{Name: "Dishant"}
	fmt.Println(user.Greet())

	user.Rename("Dishant Patel")
	fmt.Println(user.Greet())

	// recover example
	safeDivide(10, 0)
	safeDivide(10, 2) // control flow resumes as panic is recovered
}
