package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Akashkarmokar/go_rest_api/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	logger.Info("Server is starting on port : 8080")
	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Failed to start server:", err)
		logger.Info("failed to start server", "error", err)
	}
}
