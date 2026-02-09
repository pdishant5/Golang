package main

import (
	"http-demo/internal/app"
	"log"
)

func main() {
	server := app.NewServer()

	log.Println("Starting server on :8080")
	log.Fatal(server.ListenAndServe())
}
