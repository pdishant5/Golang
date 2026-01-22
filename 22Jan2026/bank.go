package main

import "fmt"

// method 1: using closure to create account operations
func createAccount(initialBalance int) func(action string, amount int) int {
	balance := initialBalance

	return func(action string, amount int) int {
		switch action {
		case "deposit":
			balance += amount
		case "withdraw":
			if amount <= balance {
				balance -= amount
			} else {
				fmt.Println("Insufficient balance")
			}
		}
		return balance
	}
}

// method 2: uses closures but cleaner design
type AccountAction func(int) int

func newAccount(balance int) (deposit AccountAction, withdraw AccountAction) {
	return func(amount int) int {
			balance += amount
			return balance
		},
		func(amount int) int {
			if amount <= balance {
				balance -= amount
			}
			return balance
		}
}

func main() {
	account1 := createAccount(1000)
	account2 := createAccount(1000)

	account1("deposit", 500)              // 1500
	account2("deposit", 800)              // 1800
	balance1 := account1("withdraw", 800) // 700
	balance2 := account2("withdraw", 500) // 1300

	fmt.Println("Account 1 Balance:", balance1) // 700
	fmt.Println("Account 2 Balance:", balance2) // 1300

	deposit1, withdraw1 := newAccount(2000)
	deposit2, withdraw2 := newAccount(3000)

	deposit1(1000)              // 3000
	withdraw1(500)              // 2500
	balanceA := withdraw1(1000) // 1500
	deposit2(2000)              // 5000
	balanceB := withdraw2(4000) // 1000

	fmt.Println("Account A Balance:", balanceA) // 1500
	fmt.Println("Account B Balance:", balanceB) // 1000
}
