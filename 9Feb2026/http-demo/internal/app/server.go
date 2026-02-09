package app

import (
	"http-demo/internal/router"
	"net/http"
	"time"
)

func NewServer() *http.Server {
	mux := router.NewRouter()

	return &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
