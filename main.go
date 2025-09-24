package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)

	slog.Info("Hello world")
	slog.Debug("Hello debug")
	slog.Error("HEllo error")

	slog.Info("Hello Info", "key1", "value1", slog.Int("Key2", 1))
	slog.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"info message",
		slog.String("key2", "2322"),
	)

	slog.Info("Group Message", slog.Group("request", slog.String("Key2", "value2")))
	// fmt.Println("Hello")

	child_logger := logger.WithGroup("Info").With(slog.String("Group Message", "My Custom Group Message"))
	child_logger.Info("My Custom Message")

}
