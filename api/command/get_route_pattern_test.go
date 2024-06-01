package command_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/command"
)

func TestGetRoutePattern(t *testing.T) {
	t.Run("No parts.", func(t *testing.T) {
		actualPattern := command.GetRoutePattern()
		expectedPattern := "POST /"

		assert.Equal(t, expectedPattern, actualPattern, fmt.Sprintf("Expected pattern to be '%s', but got '%s'", expectedPattern, actualPattern))
	})

	t.Run("Single part.", func(t *testing.T) {
		actualPattern := command.GetRoutePattern("foo")
		expectedPattern := "POST /foo"

		assert.Equal(t, expectedPattern, actualPattern, fmt.Sprintf("Expected pattern to be '%s', but got '%s'", expectedPattern, actualPattern))
	})

	t.Run("Multiple parts.", func(t *testing.T) {
		actualPattern := command.GetRoutePattern("foo", "bar")
		expectedPattern := "POST /foo/bar"

		assert.Equal(t, expectedPattern, actualPattern, fmt.Sprintf("Expected pattern to be '%s', but got '%s'", expectedPattern, actualPattern))
	})
}
