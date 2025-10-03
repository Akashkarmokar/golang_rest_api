package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Akashkarmokar/go_rest_api/internal/logger"
	"github.com/Akashkarmokar/go_rest_api/internal/router"
)

func main() {
	customLogger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))
	r := router.New()
	wrappedRouter := logger.AddLoggerMid(customLogger, logger.LoggerMid(r))
	customLogger.Info("Server is starting on port : 8080")
	if err := http.ListenAndServe(":8080", wrappedRouter); err != nil {
		log.Fatal("Failed to start server:", err)
		customLogger.Info("failed to start server", "error", err)
	}
}
