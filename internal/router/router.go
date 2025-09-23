package router

import (
	"net/http"

	"github.com/Akashkarmokar/go_rest_api/internal/handler"
)

func New() *http.ServeMux {
	r := http.NewServeMux()

	// Create news route.
	r.HandleFunc("POST /news", handler.PostNews())

	// Get all news
	r.HandleFunc("GET /news", handler.GetAllNews())

	// Get news by ID
	r.HandleFunc("GET /news/{news_id}", handler.GetNewsByID())

	// Update news by ID
	r.HandleFunc("PUT /news/{news_id}", handler.UpdateNewsByID())

	// Delete news by ID.
	r.HandleFunc("DELETE /news/{news_id}", handler.DeleteNewsByID())

	return r
}
