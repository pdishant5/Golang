package main

import (
	"encoding/json"
	"fmt"
)

func iterate(slice []int) {
	for _, v := range slice {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
}

func main() {
	var s1 []int // nil slice
	fmt.Println("Slice s1:", s1)
	fmt.Printf("Nil slice: %v; Length: %d; Cap: %d\n", s1 == nil, len(s1), cap(s1))

	s2 := []int{} // empty slice
	fmt.Println("Slice s2:", s2)
	fmt.Printf("Nil slice: %v; Length: %d; Cap: %d\n", s2 == nil, len(s2), cap(s2))

	// safe to iterate over
	iterate(s1)
	iterate(s2)

	// different behavior in converting to JSON
	data1, _ := json.Marshal(s1)
	fmt.Println("JSON of s1:", string(data1))
	data2, _ := json.Marshal(s2)
	fmt.Println("JSON of s2:", string(data2))

	// safe to append
	s1 = append(s1, 1)
	fmt.Println("\nSlice s1:", s1)
	fmt.Printf("Nil slice: %v; Length: %d; Cap: %d\n", s1 == nil, len(s1), cap(s1))

	s2 = append(s2, 1)
	fmt.Println("Slice s2:", s2)
	fmt.Printf("Nil slice: %v; Length: %d; Cap: %d\n", s2 == nil, len(s2), cap(s2))

	var i interface{}
	x := 10.0
	i = x
	// v := i.(int) // panic - not matches the type
	v, ok := i.(int) // safe - zero value in 'v' if type not matched

	fmt.Printf("Value of v: %v; Ok: %v\n", v, ok)

	var s Speaker
	p := Person{}
	s = p
	p.Speak()
	s.Speak()

	var user IUser
	// value has methods that has value receiver
	// u := User{ID: 1}
	// user = u // error assigning to IUser interface variable

	// pointer has methods with both value and pointer receiver
	u := &User{ID: 1}
	user = u // no error in assigning to IUser interface variable
	user.SetAge(20)

	fmt.Println("ID:", u.ID, "Age:", u.Age)

	var ch chan int // nil channel
	fmt.Println(ch)

	// ch <- 10 // blocks - deadlock
	// fmt.Println("Unreachable!")

	// msg := <-ch // blocks - deadlock
	// fmt.Println("Unreachable!", msg)

	// close(ch) // panic - close of nil channel

	// fmt.Println(User{"Alice"})

	// var wg sync.WaitGroup
	// wg.Add(4)
	// go func1(&wg)
	// go func2(&wg)
	// go func2(&wg)
	// go func2(&wg)

	// // time.Sleep(time.Second)
	// wg.Wait()
}

// func func1(wg *sync.WaitGroup) {
// 	wg.Done()
// 	wg.Done()
// 	wg.Done()
// 	wg.Done()
// }

// func func2(wg *sync.WaitGroup) {
// 	// wg.Done()
// 	fmt.Println("Hello!")
// }

// type User struct {
// 	Name string
// }

// func (u User) String() string {
// 	return "User: " + u.Name
// }

type IUser interface {
	GetID() int
	SetAge(int)
}

type User struct {
	ID  int
	Age int
}

func (u User) GetID() int {
	return u.ID
}

func (u *User) SetAge(age int) {
	u.Age = age
}

type Speaker interface {
	Speak()
}

type Person struct{}

func (p Person) Speak() {
	fmt.Println("Person is spreaking!")
}
