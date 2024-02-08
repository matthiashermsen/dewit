package logging

import (
	"log/slog"
	"os"
)

func GetLogger(logLevel slog.Level) *slog.Logger {
	logHandler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})

	return slog.New(logHandler)
}
