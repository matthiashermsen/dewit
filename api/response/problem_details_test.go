package response_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/response"
)

func TestNewProblemDetails(t *testing.T) {
	expectedType := "DETAIL"
	expectedStatus := http.StatusBadRequest
	expectedTitle := "made-up"
	expectedDetail := "error message"

	problemDetails := response.NewProblemDetails(expectedType, expectedStatus, expectedTitle, expectedDetail)

	assert.Equal(t, expectedType, problemDetails.Type, fmt.Sprintf("Expected type to be '%s', but got '%s'", expectedType, problemDetails.Type))
	assert.Equal(t, expectedStatus, problemDetails.Status, fmt.Sprintf("Expected status to be '%d', but got '%d'", expectedStatus, problemDetails.Status))
	assert.Equal(t, expectedTitle, problemDetails.Title, fmt.Sprintf("Expected title to be '%s', but got '%s'", expectedTitle, problemDetails.Title))
	assert.Equal(t, expectedDetail, problemDetails.Detail, fmt.Sprintf("Expected detail to be '%s', but got '%s'", expectedDetail, problemDetails.Detail))
}
