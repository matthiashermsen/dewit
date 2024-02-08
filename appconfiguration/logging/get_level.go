package logging

import (
	"log/slog"
	"strings"

	"github.com/spf13/viper"
)

func GetLevel() (slog.Level, error) {
	rawLevel := viper.GetString(LevelKey)
	rawLowerCaseLevel := strings.ToLower(rawLevel)

	switch rawLowerCaseLevel {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	case "":
		return DefaultLevel, nil
	default:
		return DefaultLevel, NewLogLevelInvalidError(rawLevel)
	}
}
