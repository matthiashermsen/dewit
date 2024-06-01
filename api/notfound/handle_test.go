package notfound_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/notfound"
	"github.com/matthiashermsen/dewit/api/response"
)

func TestHandle(t *testing.T) {
	recorder := httptest.NewRecorder()
	routePath := "/made-up"
	method := "GET"
	request, err := http.NewRequest(method, routePath, nil)

	assert.NoError(t, err)

	handler := notfound.Handle(slog.New(slog.Default().Handler()))

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusNotFound, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusNotFound, recorder.Code))

	var actualResponseBody response.ProblemDetails
	err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

	assert.NoError(t, err)

	expectedResponseDetails := response.ProblemDetails{
		Type:   "RESOURCE_MISSING",
		Status: http.StatusNotFound,
		Title:  "Not found",
		Detail: fmt.Sprintf("Could not find %s '%s'.", method, routePath),
	}

	assert.Equal(t, expectedResponseDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedResponseDetails, actualResponseBody))
}
