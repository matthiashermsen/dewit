package response_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/response"
)

func TestWriteInternalErrorProblemDetails(t *testing.T) {
	recorder := httptest.NewRecorder()

	response.WriteInternalErrorProblemDetails(recorder, slog.New(slog.Default().Handler()))

	var actualResponseBody response.ProblemDetails
	err := json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

	assert.NoError(t, err)

	expectedProblemDetails := response.NewProblemDetails("INTERNAL_ERROR", http.StatusInternalServerError, "Internal error", "The server encountered an unexpected condition that prevented it from fulfilling the request.")
	assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))

	expectedContentType := "application/problem+json"
	actualContentType := recorder.Header().Get("Content-Type")

	assert.Equal(t, expectedContentType, actualContentType, fmt.Sprintf("Expected content type to be '%s', but got '%s'", expectedContentType, actualContentType))
}
