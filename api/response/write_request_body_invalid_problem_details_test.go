package response_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/response"
)

func TestWriteRequestBodyInvalidProblemDetails(t *testing.T) {
	recorder := httptest.NewRecorder()

	detailErr := errors.New("made-up")

	expectedProblemDetails := response.NewProblemDetails("REQUEST_BODY_INVALID", http.StatusBadRequest, "Request body invalid", detailErr.Error())

	response.WriteRequestBodyInvalidProblemDetails(recorder, detailErr, slog.New(slog.Default().Handler()))

	var actualResponseBody response.ProblemDetails
	err := json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

	assert.NoError(t, err)
	assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))

	expectedContentType := "application/problem+json"
	actualContentType := recorder.Header().Get("Content-Type")

	assert.Equal(t, expectedContentType, actualContentType, fmt.Sprintf("Expected content type to be '%s', but got '%s'", expectedContentType, actualContentType))
}
