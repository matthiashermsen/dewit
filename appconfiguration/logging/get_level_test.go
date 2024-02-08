package logging_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/appconfiguration/logging"
)

func TestGetLevel(t *testing.T) {
	t.Run("Default Level", func(t *testing.T) {
		actualLevel, err := logging.GetLevel()

		assert.NoError(t, err)

		assert.Equal(t, logging.DefaultLevel, actualLevel, fmt.Sprintf("Expected level to be '%s', but got '%s'", logging.DefaultLevel, actualLevel))
	})

	t.Run("Custom Level", func(t *testing.T) {
		t.Run("Debug level", func(t *testing.T) {
			viper.Set(logging.LevelKey, "debug")
			actualLevel, err := logging.GetLevel()

			assert.NoError(t, err)

			assert.Equal(t, slog.LevelDebug, actualLevel, fmt.Sprintf("Expected level to be '%s', but got '%s'", slog.LevelDebug, actualLevel))
		})

		t.Run("Info level", func(t *testing.T) {
			viper.Set(logging.LevelKey, "info")
			actualLevel, err := logging.GetLevel()

			assert.NoError(t, err)

			assert.Equal(t, slog.LevelInfo, actualLevel, fmt.Sprintf("Expected level to be '%s', but got '%s'", slog.LevelInfo, actualLevel))
		})

		t.Run("Warn level", func(t *testing.T) {
			viper.Set(logging.LevelKey, "warn")
			actualLevel, err := logging.GetLevel()

			assert.NoError(t, err)

			assert.Equal(t, slog.LevelWarn, actualLevel, fmt.Sprintf("Expected level to be '%s', but got '%s'", slog.LevelWarn, actualLevel))
		})

		t.Run("Error level", func(t *testing.T) {
			viper.Set(logging.LevelKey, "error")
			actualLevel, err := logging.GetLevel()

			assert.NoError(t, err)

			assert.Equal(t, slog.LevelError, actualLevel, fmt.Sprintf("Expected level to be '%s', but got '%s'", slog.LevelError, actualLevel))
		})

		t.Run("Empty level", func(t *testing.T) {
			viper.Set(logging.LevelKey, "")
			actualLevel, err := logging.GetLevel()

			assert.NoError(t, err)

			assert.Equal(t, logging.DefaultLevel, actualLevel, fmt.Sprintf("Expected level to be '%s', but got '%s'", logging.DefaultLevel, actualLevel))
		})

		t.Run("Invalid level", func(t *testing.T) {
			viper.Set(logging.LevelKey, "made-up")
			_, err := logging.GetLevel()

			assert.Error(t, err, "Expected log level to be invalid")
		})
	})
}
