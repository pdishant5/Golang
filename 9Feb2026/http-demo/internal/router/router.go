package router

import (
	"net/http"

	"http-demo/internal/handler"
	"http-demo/internal/middleware"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	// handlers
	mux.HandleFunc("/health", handler.Health)

	mux.Handle("/users", middleware.Logging(
		http.HandlerFunc(handler.Users),
	))

	mux.Handle("/users/", middleware.Logging(
		http.HandlerFunc(handler.UserByID),
	))

	return mux
}
