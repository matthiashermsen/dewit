package response_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/response"
)

func TestJSONParseErrorProblemDetails(t *testing.T) {
	expectedProblemDetails := response.
		NewProblemDetails().
		WithType("JSON_PARSE_ERROR").
		WithStatus(http.StatusBadRequest).
		WithTitle("Could not parse JSON request body").
		WithDetail("The request body is invalid.")

	actualProblemDetails := response.GetJSONParseErrorProblemDetails()

	assert.Equal(t, expectedProblemDetails, actualProblemDetails, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualProblemDetails))
}
