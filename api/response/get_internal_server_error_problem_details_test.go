package response_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/response"
)

func TestGetInternalServerErrorProblemDetails(t *testing.T) {
	expectedProblemDetails := response.
		NewProblemDetails().
		WithType("INTERNAL_ERROR").
		WithStatus(http.StatusInternalServerError).
		WithTitle("Internal error").
		WithDetail("The server encountered an unexpected condition that prevented it from fulfilling the request.")

	actualProblemDetails := response.GetInternalServerErrorProblemDetails()

	assert.Equal(t, expectedProblemDetails, actualProblemDetails, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualProblemDetails))
}
