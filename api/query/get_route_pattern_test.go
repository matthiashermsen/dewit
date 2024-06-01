package query_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/query"
)

func TestGetRoutePattern(t *testing.T) {
	t.Run("No parts.", func(t *testing.T) {
		actualPattern := query.GetRoutePattern()
		expectedPattern := "GET /"

		assert.Equal(t, expectedPattern, actualPattern, fmt.Sprintf("Expected pattern to be '%s', but got '%s'", expectedPattern, actualPattern))
	})

	t.Run("Single part.", func(t *testing.T) {
		actualPattern := query.GetRoutePattern("foo")
		expectedPattern := "GET /foo"

		assert.Equal(t, expectedPattern, actualPattern, fmt.Sprintf("Expected pattern to be '%s', but got '%s'", expectedPattern, actualPattern))
	})

	t.Run("Multiple parts.", func(t *testing.T) {
		actualPattern := query.GetRoutePattern("foo", "bar")
		expectedPattern := "GET /foo/bar"

		assert.Equal(t, expectedPattern, actualPattern, fmt.Sprintf("Expected pattern to be '%s', but got '%s'", expectedPattern, actualPattern))
	})
}
