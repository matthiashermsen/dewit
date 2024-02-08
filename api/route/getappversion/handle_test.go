package getappversion_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/matthiashermsen/dewit/api/response"
	"github.com/matthiashermsen/dewit/api/route/getappversion"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {
	t.Run("App version is available", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/app-version", nil)

		assert.NoError(t, err)

		expectedAppVersion := "1.2.3"
		handler := getappversion.Handle(slog.New(slog.Default().Handler()), expectedAppVersion)

		handler.ServeHTTP(recorder, request)

		var actualResponseBody string
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		assert.Equal(t, expectedAppVersion, actualResponseBody, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedAppVersion, actualResponseBody))
	})

	t.Run("App version is unavailable", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/app-version", nil)

		assert.NoError(t, err)

		handler := getappversion.Handle(slog.New(slog.Default().Handler()), "")

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusNotFound, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		expectedResponseDetails := response.
			NewProblemDetails().
			WithType("APP_VERSION_UNAVAILABLE").
			WithStatus(http.StatusNotFound).
			WithTitle("App version is unavailable").
			WithDetail("Current app version is not available.")

		assert.Equal(t, expectedResponseDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedResponseDetails, actualResponseBody))
	})
}
