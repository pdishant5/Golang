package main

import (
	"fmt"
	// "price-calculator/cmdmanager"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, taxRate := range taxRates {
		inputFilePath := "prices.txt"
		outputFilePath := fmt.Sprintf("result_%.0f.json", taxRate*100)

		fm := filemanager.New(inputFilePath, outputFilePath)
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Error processing the job!")
			fmt.Println("Error:", err)
		}
	}
}
