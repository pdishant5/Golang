package handler

import (
	"net/http"
	"strings"

	"http-demo/pkg/response"
)

// Query parameters example: /users?limit=10
func Users(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")

	response.JSON(w, http.StatusOK, map[string]string{
		"message": "users list",
		"limit":   limit,
	})
}

// Path parameter example: /users/123
func UserByID(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")

	response.JSON(w, http.StatusOK, map[string]string{
		"user_id": id,
	})
}
