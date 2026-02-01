package main

import (
	"fmt"
	// "price-calculator/cmdmanager"
	"price-calculator/filemanager"
	"price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		inputFilePath := "prices.txt"
		outputFilePath := fmt.Sprintf("result_%.0f.json", taxRate*100)

		fm := filemanager.New(inputFilePath, outputFilePath)
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// err := priceJob.Process()
		go priceJob.Process(doneChans[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Error processing the job!")
		// 	fmt.Println("Error:", err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println("Error:", err)
			}
		case <-doneChans[index]:
			fmt.Println("Done!")
		}
	}

	// for _, errorChan := range errorChans {
	// 	<-errorChan
	// } // this doesn't work as expected - it waits for the message in the error channel, and then the program crashes

	// for _, doneChan := range doneChans {
	// 	<-doneChan
	// }
}
