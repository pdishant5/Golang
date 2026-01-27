package main

import (
	"interfaces/person"
	"interfaces/user"
)

// type Person struct {
// 	name string
// 	age  int
// }

// func (p Person) GetName() string {
// 	return p.name
// }

// func (p Person) getAge() int {
// 	return p.age
// }

// type User struct {
// 	name string
// 	age  int
// }

// func (u User) GetName() string {
// 	return u.name
// }

// func (u User) getAge() int {
// 	return u.age
// }

type Person interface {
	GetName() string
	getAge() int
}

type User interface {
	GetName() string
	getAge() int
}

func main() {
	var P Person
	var U User

	var p person.Person
	var u user.User

	P = U // valid assignment since User implements Person
	p = u // invalid assignment since both interfaces have unexported methods - they are distinct

	// just checking swap of two variables
	// a := 10
	// b := 5
	// a, b = b, a // swap values of a and b
	// fmt.Println("a:", a, "b:", b)
}
