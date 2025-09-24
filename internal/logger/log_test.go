package logger_test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/Akashkarmokar/go_rest_api/internal/logger"
)

func Test_ContextWithLogger(t *testing.T) {
	testCases := []struct {
		name            string
		ctx             context.Context
		logger          *slog.Logger
		doesLoggerExist bool
	}{
		{
			name:            "returns ctx without logger",
			ctx:             context.Background(),
			logger:          nil,   // Explicitly typed, although not required
			doesLoggerExist: false, // Explicitly typed here
		},
		{
			name: "return ctx as it is",
			ctx: context.WithValue(
				context.Background(),
				logger.CtxKey{},
				slog.New(
					slog.NewJSONHandler(
						os.Stdout,
						&slog.HandlerOptions{
							AddSource: true,
							Level:     slog.LevelDebug,
						},
					),
				),
			),
			doesLoggerExist: true,
		},
		{
			name: "inject logger",
			ctx:  context.Background(),
			logger: slog.New(
				slog.NewJSONHandler(
					os.Stdout,
					&slog.HandlerOptions{
						AddSource: true,
						Level:     slog.LevelDebug,
					},
				),
			),
			doesLoggerExist: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := logger.CtxWithLogger(tc.ctx, tc.logger)

			_, ok := ctx.Value(logger.CtxKey{}).(*slog.Logger)
			if tc.doesLoggerExist != ok {
				t.Errorf("Expected :%v got: %v", false, ok)
			}
		})
	}
}

func Test_FromContext(t *testing.T) {
	testCases := []struct {
		name     string
		ctx      context.Context
		expected bool
	}{
		{
			name: "logger exists",
			ctx: context.WithValue(
				context.Background(),
				logger.CtxKey{},
				slog.New(
					slog.NewJSONHandler(
						os.Stdout,
						&slog.HandlerOptions{
							AddSource: true,
							Level:     slog.LevelDebug,
						},
					),
				),
			),
		},
		{
			name:     "new logger returned",
			ctx:      context.Background(),
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger := logger.FromContext(tc.ctx)
			if tc.expected && logger == nil {
				t.Errorf("Expected %v, got: %v", tc.expected, logger)
			}
		})
	}
}
