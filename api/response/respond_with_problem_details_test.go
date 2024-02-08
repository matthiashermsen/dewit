package response_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiashermsen/dewit/api/response"

	"github.com/stretchr/testify/assert"
)

func TestRespondWithProblemDetails(t *testing.T) {
	recorder := httptest.NewRecorder()

	expectedProblemDetails := response.NewProblemDetails().
		WithType("NOT_FOUND").
		WithStatus(http.StatusNotFound).
		WithTitle("Not found").
		WithDetail("Could not find thing.")

	response.RespondWithProblemDetails(expectedProblemDetails, recorder, slog.New(slog.Default().Handler()))

	var actualResponseBody response.ProblemDetails
	err := json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

	assert.NoError(t, err)
	assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
}
