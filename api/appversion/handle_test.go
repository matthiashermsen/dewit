package appversion_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/appversion"
	"github.com/matthiashermsen/dewit/api/response"
)

func TestHandle(t *testing.T) {
	t.Run("App version is available", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/app-version", nil)

		assert.NoError(t, err)

		expectedAppVersion := "1.2.3"
		handler := appversion.Handle(expectedAppVersion, slog.New(slog.Default().Handler()))

		handler.ServeHTTP(recorder, request)

		var actualResponseBody string
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)
		assert.Equal(t, expectedAppVersion, actualResponseBody, fmt.Sprintf("Expected response body '%s', but got '%s'", expectedAppVersion, actualResponseBody))
	})

	t.Run("App version is unavailable", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		httpRequest, err := http.NewRequest("GET", "/app-version", nil)

		assert.NoError(t, err)

		handler := appversion.Handle("", slog.New(slog.Default().Handler()))

		handler.ServeHTTP(recorder, httpRequest)

		assert.Equal(t, http.StatusNotFound, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusNotFound, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		expectedResponseDetails := response.ProblemDetails{
			Type:   "APP_VERSION_UNAVAILABLE",
			Status: http.StatusNotFound,
			Title:  "App version is unavailable",
			Detail: "Current app version is not available.",
		}

		assert.Equal(t, expectedResponseDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedResponseDetails, actualResponseBody))
	})
}
