package log

import (
	"log/slog"
	"os"
)

func New(logLevel slog.Level) *slog.Logger {
	logHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})

	return slog.New(logHandler)
}
