package response_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/response"
)

func TestNewProblemDetails(t *testing.T) {
	rd := response.NewProblemDetails()

	assert.Equal(t, "", rd.Type, fmt.Sprintf("Expected type to be '', but got '%s'", rd.Type))
	assert.Equal(t, 0, rd.Status, fmt.Sprintf("Expected status to be 0, but got '%d'", rd.Status))
	assert.Equal(t, "", rd.Title, fmt.Sprintf("Expected title to be '', but got '%s'", rd.Title))
	assert.Equal(t, "", rd.Detail, fmt.Sprintf("Expected detail to be '', but got '%s'", rd.Detail))
}

func TestWithType(t *testing.T) {
	rdType := "made-up"
	rd := response.NewProblemDetails().WithType(rdType)

	assert.Equal(t, rdType, rd.Type, fmt.Sprintf("Expected type to be '%s', but got '%s'", rdType, rd.Type))
}

func TestWithStatus(t *testing.T) {
	rdStatus := 200
	rd := response.NewProblemDetails().WithStatus(rdStatus)

	assert.Equal(t, rdStatus, rd.Status, fmt.Sprintf("Expected status to be '%d', but got '%d'", rdStatus, rd.Status))
}

func TestWithTitle(t *testing.T) {
	rdTitle := "made-up"
	rd := response.NewProblemDetails().WithTitle(rdTitle)

	assert.Equal(t, rdTitle, rd.Title, fmt.Sprintf("Expected title to be '%s', but got '%s'", rdTitle, rd.Title))
}

func TestWithDetail(t *testing.T) {
	rdDetail := "made-up"
	rd := response.NewProblemDetails().WithDetail(rdDetail)

	assert.Equal(t, rdDetail, rd.Detail, fmt.Sprintf("Expected detail to be '%s', but got '%s'", rdDetail, rd.Detail))
}

func TestWithInternalError(t *testing.T) {
	rd := response.NewProblemDetails().WithInternalError()

	assert.Equal(t, "INTERNAL_ERROR", rd.Type, fmt.Sprintf("Expected type to be 'INTERNAL_ERROR', but got '%s'", rd.Type))
	assert.Equal(t, http.StatusInternalServerError, rd.Status, fmt.Sprintf("Expected status to be %d, but got '%d'", http.StatusInternalServerError, rd.Status))

	expectedTitle := "Internal error"
	assert.Equal(t, expectedTitle, rd.Title, fmt.Sprintf("Expected title to be '%s', but got '%s'", expectedTitle, rd.Title))

	expectedDetail := "The server encountered an unexpected condition that prevented it from fulfilling the request."
	assert.Equal(t, expectedDetail, rd.Detail, fmt.Sprintf("Expected detail to be '%s', but got '%s'", expectedDetail, rd.Detail))
}
