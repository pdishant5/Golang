package main

import (
	"app/fileops"
	"fmt"
)

const accountBalanceFile = "balance.txt"

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile, 0.0)

	if err != nil {
		accountBalance = 0.0
		fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
	}

	fmt.Println("Welcome to the Banking App!")
	// switch-case
	for {
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
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
			fileops.WriteFloatToFile(accountBalanceFile, accountBalance)
		default:
			fmt.Println("Thank you! Visit again!")
			return
		}
	}
}
