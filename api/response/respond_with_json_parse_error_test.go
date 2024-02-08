package response_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http/httptest"
	"testing"

	"github.com/matthiashermsen/dewit/api/response"

	"github.com/stretchr/testify/assert"
)

func TestRespondWithJSONParseError(t *testing.T) {
	recorder := httptest.NewRecorder()

	response.RespondWithJSONParseError(recorder, slog.New(slog.Default().Handler()))

	var actualResponseBody response.ProblemDetails
	err := json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

	expectedProblemDetails := response.GetJSONParseErrorProblemDetails()

	assert.NoError(t, err)
	assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
}
