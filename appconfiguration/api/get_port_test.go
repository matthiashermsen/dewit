package api_test

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/appconfiguration/api"
)

func TestGetPort(t *testing.T) {
	t.Run("Default Port", func(t *testing.T) {
		actualPort := api.GetPort()

		assert.Equal(t, api.DefaultPort, actualPort, fmt.Sprintf("Expected port to be '%d', but got '%d'", api.DefaultPort, actualPort))
	})

	t.Run("Custom Port", func(t *testing.T) {
		var expectedPort uint16 = 1234

		viper.Set(api.PortKey, expectedPort)

		actualPort := api.GetPort()

		assert.Equal(t, expectedPort, actualPort, fmt.Sprintf("Expected port to be '%d', but got '%d'", expectedPort, actualPort))
	})
}
