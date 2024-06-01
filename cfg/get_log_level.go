package cfg

import (
	"log/slog"
	"os"
)

const (
	LogLevelKey     = "LOG_LEVEL"
	DefaultLogLevel = slog.LevelInfo
)

func GetLogLevel() (slog.Level, error) {
	rawLogLevel, hasLogLevel := os.LookupEnv(LogLevelKey)

	if !hasLogLevel || rawLogLevel == "" {
		return DefaultLogLevel, nil
	}

	var level slog.Level
	err := level.UnmarshalText([]byte(rawLogLevel))

	return level, err
}
