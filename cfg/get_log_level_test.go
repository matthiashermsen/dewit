package cfg_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/cfg"
)

func TestLogLevelKey(t *testing.T) {
	expectedLogLevelKey := "LOG_LEVEL"
	assert.Equal(t, expectedLogLevelKey, cfg.LogLevelKey, fmt.Sprintf("Expected log level key to be '%s', but got '%s'", expectedLogLevelKey, cfg.LogLevelKey))
}

func TestDefaultLogLevel(t *testing.T) {
	expectedDefaultLogLevel := slog.LevelInfo
	assert.Equal(t, expectedDefaultLogLevel, cfg.DefaultLogLevel, fmt.Sprintf("Expected default log level be '%s', but got '%s'", expectedDefaultLogLevel, cfg.DefaultLogLevel))
}

func TestGetLogLevel(t *testing.T) {
	t.Run("Returns default log level if missing.", func(t *testing.T) {
		t.Setenv(cfg.LogLevelKey, "")

		actualLogLevel, err := cfg.GetLogLevel()

		assert.NoError(t, err)
		assert.Equal(t, cfg.DefaultLogLevel, actualLogLevel, fmt.Sprintf("Expected log level to be '%s', but got '%s'", cfg.DefaultLogLevel, actualLogLevel))
	})

	t.Run("Returns error if log level is invalid.", func(t *testing.T) {
		t.Setenv(cfg.LogLevelKey, "MADEUP")

		_, err := cfg.GetLogLevel()

		assert.Error(t, err, "Expected log level to be invalid")
	})

	t.Run("Returns log level.", func(t *testing.T) {
		expectedLogLevel := slog.LevelWarn
		t.Setenv(cfg.LogLevelKey, expectedLogLevel.String())

		actualLogLevel, err := cfg.GetLogLevel()

		assert.NoError(t, err)
		assert.Equal(t, expectedLogLevel, actualLogLevel, fmt.Sprintf("Expected log level to be '%s', but got '%s'", expectedLogLevel, actualLogLevel))
	})
}
