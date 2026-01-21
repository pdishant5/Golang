package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter your annual revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your annual expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100)

	ratio := earningsBeforeTax / earningsAfterTax

	fmt.Println("Your earnings before tax (EBT) is:", earningsBeforeTax)
	fmt.Println("Your earnings after tax (profit) is:", earningsAfterTax)
	fmt.Println("Your EBT to profit ratio is:", ratio)
}
