package main

import (
	"fmt"
	"unsafe"
)

type Bad struct {
	A int32
	B int64
	C byte
}

// memory alignment optimization - ordering fields by size
type Good struct {
	B int64
	A int32
	C byte
}

type User struct {
	ID   int
	Name string
	Age  int
}

// value receiver method
func (u User) IsAdult() bool {
	return u.Age >= 18
}

// pointer receiver method
func (u *User) SetAge(age int) {
	u.Age = age
}

// u.SetAge(20) --> SetAge(&u, 20)
// methods are just functions with a special receiver argument

// u *User --> u.Age and (*u).Age are equivalent

func main() {
	var bad Bad
	var good Good

	fmt.Println("Size of Bad struct:", unsafe.Sizeof(bad))   // 24 bytes due to padding
	fmt.Println("Size of Good struct:", unsafe.Sizeof(good)) // 16 bytes without padding

	var u User // zero-value initialization
	fmt.Println("User ID:", u.ID, "Name:", u.Name, "Age:", u.Age)

	u1 := User{1, "Dishant", 20} // composite literal initialization
	fmt.Println("User ID:", u1.ID, "Name:", u1.Name, "Age:", u1.Age)

	u2 := User{ID: 1, Name: "Dishant", Age: 20} // initialization with named fields
	fmt.Println("User ID:", u2.ID, "Name:", u2.Name, "Age:", u2.Age)

	var u3 = User{Name: "Dishant"} // partial initialization, other fields get zero values
	fmt.Println("User ID:", u3.ID, "Name:", u3.Name, "Age:", u3.Age)

	point := struct {
		x int
		y int
	}{10, 20} // anonymous struct

	fmt.Println("Point coordinates:", "X:", point.x, "Y:", point.y)

	u4 := User{ID: 1}
	u5 := u4
	u5.ID = 2
	fmt.Println("u4 ID:", u4.ID, "u5 ID", u5.ID) // 1, 2 - value is copied, u1 remains unchanged

	u6 := &u4
	u6.ID = 3
	fmt.Println("u4 ID:", u4.ID, "u6 ID", u6.ID) // 3, 3 - both reflect the same underlying data
}
