package main

import "fmt"

func main() {
	var productNames [5]string = [5]string{"Thumbs up", "Cookie", "Biscuit"}
	prices := [5]float64{19.99, 29.99, 4.99, 49.99, 9.99}
	fmt.Println("Array of prices:", prices)

	productNames[3], productNames[4] = "Dairy Milk Silk", "5 Star"
	fmt.Println("Array of product names:", productNames)

	for i := 0; i < len(prices); i++ {
		fmt.Printf("Price of %s: %.2f\n", productNames[i], prices[i])
	}

	featuredPrices := prices[1:4] // slicing the array from index 1 to 3
	fmt.Println("Featured Prices:", featuredPrices)

	highlightedPrices := featuredPrices[:2] // further slicing the featuredPrices slice
	fmt.Println("Highlighted Prices:", highlightedPrices)

	featuredPrices[0] = 24.99 // modifying the slice also modifies the original array
	fmt.Println("Modified Featured Prices:", featuredPrices)
	fmt.Println("Updated Prices Array:", prices)

	// creating a slice using make
	newPrices := make([]float64, 3)
	newPrices[0] = 14.99
	newPrices[1] = 34.99
	newPrices[2] = 7.99
	fmt.Println("New Prices Slice:", newPrices)
	fmt.Println("Size of the Slice:", len(newPrices))     // 3
	fmt.Println("Capacity of the Slice:", cap(newPrices)) // 3

	// appending to a slice
	newPrices = append(newPrices, 44.99, 54.99)
	fmt.Println("Appended New Slice:", newPrices)
	fmt.Println("Size of the Slice:", len(newPrices))     // 5
	fmt.Println("Capacity of the Slice:", cap(newPrices)) // 6

	newPrices = append(newPrices, 7, 8, 9, 10, 12, 8, 8)
	fmt.Println("Appended New Slice:", newPrices)
	fmt.Println("Size of the Slice:", len(newPrices))     // 12
	fmt.Println("Capacity of the Slice:", cap(newPrices)) // 12
	newPrices = append(newPrices, 64.99)
	fmt.Println("Appended New Slice:", newPrices)
	fmt.Println("Size of the Slice:", len(newPrices))     // 13
	fmt.Println("Capacity of the Slice:", cap(newPrices)) // 24

	// copying a slice
	copiedPrices := make([]float64, 5)
	// copiedPrices := newPrices[1:5] // alternative way to create a slice of first 5 elements, but capacity will be of the original slice - shallow copy
	copy(copiedPrices, newPrices[:5])
	fmt.Println("Copied Prices Slice:", copiedPrices)
	fmt.Println("Size of the Copied Slice:", len(copiedPrices))
	fmt.Println("Capacity of the Copied Slice:", cap(copiedPrices))

	// modifying the copied slice, but original is unaffected - deep copy
	copiedPrices[0] = 99.99
	fmt.Println("Modified Copied Slice:", copiedPrices)
	fmt.Println("Original Slice after modifying copy:", newPrices)

	newNewPrices := copiedPrices[1:5] // slicing the copied slice
	fmt.Println("Size of the Copied Slice:", len(newNewPrices))
	fmt.Println("Capacity of the Copied Slice:", cap(newNewPrices))
}
