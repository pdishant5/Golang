package main

import (
	"fmt"
	"struct/user"
)

// custom type
type str string

// method for custom type
func (text str) log() {
	fmt.Println(text)
}

func main() {
	userFirstName := getUserData("Enter your first name: ")
	userLastName := getUserData("Enter your last name: ")
	userBirthdate := getUserData("Enter your  birthdate (DD/MM/YYYY): ")

	// var newUser *User
	var newUser *user.User

	// user = User{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthdate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	// newUser, err := newUser(userFirstName, userLastName, userBirthdate)
	newUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	admin := user.NewAdmin("admin@gmail.com", "admin1234") // Admin struct embedding User struct

	admin.PrintUserDetails()
	admin.ClearUserName()
	admin.PrintUserDetails()

	// printUserDetails(user)
	// printUserDetails(&user)

	// newUser.printUserDetails()
	// newUser.clearUserName()
	// newUser.printUserDetails()

	newUser.PrintUserDetails()
	newUser.ClearUserName()
	newUser.PrintUserDetails()

	var message str = "User created successfully!"
	message.log()
}

// func printUserDetails(u *User) {
// 	fmt.Println("\nUser Information:")
// 	fmt.Println("First Name:", u.firstName)
// 	fmt.Println("Last Name:", u.lastName)
// 	fmt.Println("Birthdate:", u.birthdate)
// }

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)
	return value
}
