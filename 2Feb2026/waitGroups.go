package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var wg2 sync.WaitGroup

func main() {
	// wait group example
	websitesList := []string{
		"https://google.com",
		"https://gooogle.com",
		"https://github.com",
		"https://youtube.com",
		"https://keka.com",
	}

	for _, web := range websitesList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()

	// panic when counter becomes negative
	wg2.Add(1)
	go func() {
		fmt.Println("Done once!")
		defer wg2.Done()
		go func() {
			// time.Sleep(time.Millisecond)
			fmt.Println("Done twice!")
			wg2.Done() // happens later
		}()
		wg2.Wait()
	}()

	wg2.Wait()

	// You think everything is done..
	wg2.Add(1) // unsafe
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Error getting the endpoint:", endpoint)
		return
	}

	fmt.Println("Endpoint:", endpoint, ", Response:", res.StatusCode)
	// wg.Done() // doesn't run for the error case - counter never becomes 0, waited indefinitely
}
