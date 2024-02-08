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

func TestRespondWithInternalServerError(t *testing.T) {
	recorder := httptest.NewRecorder()

	response.RespondWithInternalServerError(recorder, slog.New(slog.Default().Handler()))

	var actualResponseBody response.ProblemDetails
	err := json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

	expectedProblemDetails := response.
		NewProblemDetails().
		WithType("INTERNAL_ERROR").
		WithStatus(http.StatusInternalServerError).
		WithTitle("Internal error").
		WithDetail("The server encountered an unexpected condition that prevented it from fulfilling the request.")

	assert.NoError(t, err)
	assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
}
