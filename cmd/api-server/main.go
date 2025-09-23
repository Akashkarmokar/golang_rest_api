package main

import (
	"log"
	"net/http"

	"github.com/Akashkarmokar/go_rest_api/internal/router"
)

func main() {
	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
