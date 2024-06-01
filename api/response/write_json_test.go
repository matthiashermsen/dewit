package response_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/response"
)

func TestWriteJSON(t *testing.T) {
	expectedResponse := "made-up"
	recorder := httptest.NewRecorder()

	response.WriteJSON(recorder, expectedResponse, slog.New(slog.Default().Handler()))

	var actualResponseBody string
	err := json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, actualResponseBody, fmt.Sprintf("Expected response to be '%s', but got '%#v'", expectedResponse, actualResponseBody))

	expectedContentType := "application/json"
	actualContentType := recorder.Header().Get("Content-Type")

	assert.Equal(t, expectedContentType, actualContentType, fmt.Sprintf("Expected content type to be '%s', but got '%s'", expectedContentType, actualContentType))
}
