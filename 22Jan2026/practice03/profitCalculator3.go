package main

import (
	"app/fileops"
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	var err error

	fileops.WriteFloatToFile("dummy.txt", 644.0)
input1:
	PrintOutput("Enter your annual revenue: ")
	revenue, err = TakeInput()

	if err != nil {
		fmt.Println(err)
		goto input1
	}

input2:
	PrintOutput("Enter your annual expenses: ")
	expenses, err = TakeInput()

	if err != nil {
		fmt.Println(err)
		goto input2
	}

input3:
	PrintOutput("Enter your tax rate: ")
	taxRate, err = TakeInput()

	if err != nil {
		fmt.Println(err)
		goto input3
	}

	earningsBeforeTax, earningsAfterTax, ratio := CalculateFinancials(revenue, expenses, taxRate)

	yourEBT := fmt.Sprintf("Your earnings before tax (EBT) is: %.2f\n", earningsBeforeTax)
	yourProfit := fmt.Sprintf("Your earnings after tax (profit) is: %.2f\n", earningsAfterTax)
	yourRatio := fmt.Sprintf("Your EBT to profit ratio is: %.2f\n", ratio)

	PrintOutput(yourEBT)
	PrintOutput(yourProfit)
	PrintOutput(yourRatio)

	writeFinancialsToFile(yourEBT + yourProfit + yourRatio)
}

func writeFinancialsToFile(data string) {
	os.WriteFile("financials.txt", []byte(data), 0644)
}

func PrintOutput(text string) {
	fmt.Print(text)
}

func TakeInput() (float64, error) {
	var value float64
	fmt.Scan(&value)
	if value <= 0 {
		return 0, errors.New("Invalid input: value must be positive!")
	}
	return value, nil
}

func CalculateFinancials(revenue, expenses, taxRate float64) (EBT, profit, ratio float64) {
	EBT = revenue - expenses
	profit = EBT * (1 - taxRate/100)
	ratio = EBT / profit
	return EBT, profit, ratio
}
