package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {

	var accountBalance = getBalanceFromFile()

	fmt.Println("Welcome to the Banking App!")
	// if-elseif-else ladder
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Your Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		fmt.Println("Your choice:", choice)

		if choice == 1 {
			fmt.Println("Your current balance is:", accountBalance)
		} else if choice == 2 {
			var depositAmount float64
			fmt.Print("Enter Deposit Amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount!")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! Your current balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			var withdrawAmount float64
			fmt.Print("Enter Withdraw Amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount!")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("You can't withdraw", withdrawAmount, "amount due to insufficient balance!")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! Your current balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println("Thank you! Visit again!")
			break
		}
	}

	fmt.Println("Welcome to the Banking App!")
	// switch-case
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Your Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		fmt.Println("Your choice:", choice)

		switch choice {
		case 1:
			fmt.Println("Your current balance is:", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("Enter Deposit Amount: ")
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid Amount!")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! Your current balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			var withdrawAmount float64
			fmt.Print("Enter Withdraw Amount: ")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid Amount!")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("You can't withdraw", withdrawAmount, "amount due to insufficient balance!")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! Your current balance is:", accountBalance)
			writeBalanceToFile(accountBalance)
		default:
			fmt.Println("Thank you! Visit again!")
			return
		}
	}
}
