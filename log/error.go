package log

import "log/slog"

func Error(logger *slog.Logger, msg string, err error) {
	logger.Error(msg, slog.Any("error", err))
}
