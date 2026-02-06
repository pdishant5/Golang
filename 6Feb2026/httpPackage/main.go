package main

import (
	"fmt"
	"net/http"
)

func apiHandler(res http.ResponseWriter, req *http.Request) {
	// writes "Hello, World!" into the response variable "res"
	fmt.Fprintln(res, "Hello, World!")
}

func main() {
	// in Node.js terminology - "controller function"
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/user", apiHandler)

	fmt.Println("Server started on port 8000...")
	http.ListenAndServe(":8000", nil) // ":8000" listens and serves at all interfaces
}
