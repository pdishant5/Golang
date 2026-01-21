package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	PrintOutput("Enter your annual revenue: ")
	revenue = TakeInput()

	PrintOutput("Enter your annual expenses: ")
	expenses = TakeInput()

	PrintOutput("Enter your tax rate: ")
	taxRate = TakeInput()

	earningsBeforeTax, earningsAfterTax, ratio := CalculateFinancials(revenue, expenses, taxRate)

	yourEBT := fmt.Sprintf("Your earnings before tax (EBT) is: %f\n", earningsBeforeTax)
	yourProfit := fmt.Sprintf("Your earnings after tax (profit) is: %f\n", earningsAfterTax)
	yourRatio := fmt.Sprintf("Your EBT to profit ratio is: %f\n", ratio)

	PrintOutput(yourEBT)
	PrintOutput(yourProfit)
	PrintOutput(yourRatio)
}

func PrintOutput(text string) {
	fmt.Print(text)
}

func TakeInput() (value float64) {
	fmt.Scan(&value)
	return value
}

func CalculateFinancials(revenue, expenses, taxRate float64) (EBT, profit, ratio float64) {
	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = EBT / profit
	return EBT, profit, ratio
}
