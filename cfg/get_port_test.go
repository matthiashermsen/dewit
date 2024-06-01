package cfg_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/cfg"
)

func TestPortKey(t *testing.T) {
	expectedPortKey := "PORT"
	assert.Equal(t, expectedPortKey, cfg.PortKey, fmt.Sprintf("Expected port key to be '%s', but got '%s'", expectedPortKey, cfg.PortKey))
}

func TestDefaultPort(t *testing.T) {
	expectedDefaultPort := "8080"
	assert.Equal(t, expectedDefaultPort, cfg.DefaultPort, fmt.Sprintf("Expected default port be '%s', but got '%s'", expectedDefaultPort, cfg.DefaultPort))
}

func TestGetPort(t *testing.T) {
	t.Run("Returns default port if missing.", func(t *testing.T) {
		t.Setenv(cfg.PortKey, "")

		actualPort := cfg.GetPort()

		assert.Equal(t, cfg.DefaultPort, actualPort, fmt.Sprintf("Expected port to be '%s', but got '%s'", cfg.DefaultPort, actualPort))
	})

	t.Run("Returns port.", func(t *testing.T) {
		expectedPort := "1234"
		t.Setenv(cfg.PortKey, expectedPort)

		actualPort := cfg.GetPort()

		assert.Equal(t, expectedPort, actualPort, fmt.Sprintf("Expected port to be '%s', but got '%s'", expectedPort, actualPort))
	})
}
