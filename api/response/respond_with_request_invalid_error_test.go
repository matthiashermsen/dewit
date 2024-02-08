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

func TestRespondWithRequestInvalidError(t *testing.T) {
	recorder := httptest.NewRecorder()
	expectedDetail := "The request data is invalid."

	response.RespondWithRequestInvalidError(expectedDetail, recorder, slog.New(slog.Default().Handler()))

	var actualResponseBody response.ProblemDetails
	err := json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

	expectedProblemDetails := response.
		NewProblemDetails().
		WithType("REQUEST_INVALID").
		WithStatus(http.StatusBadRequest).
		WithTitle("Invalid request").
		WithDetail(expectedDetail)

	assert.NoError(t, err)
	assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
}
