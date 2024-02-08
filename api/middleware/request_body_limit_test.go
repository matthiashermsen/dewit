package middleware_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/middleware"
	"github.com/matthiashermsen/dewit/api/response"
)

func TestRequestBodyLimit(t *testing.T) {
	t.Run("Content too large", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		maximumBytes := int64(1)
		handler := func(w http.ResponseWriter, r *http.Request) {
			t.Error("Handler should not be called since request body size is exceeding limit")
		}
		middlewareHandler := middleware.RequestBodyLimit(logger, maximumBytes, handler)

		request, err := http.NewRequest("POST", "/", strings.NewReader("LargeRequestBody"))

		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		middlewareHandler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusRequestEntityTooLarge, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusRequestEntityTooLarge, recorder.Code))

		expectedResponseBody := response.
			NewProblemDetails().
			WithType("CONTENT_TOO_LARGE").
			WithStatus(http.StatusRequestEntityTooLarge).
			WithTitle("Body size exceeding limit").
			WithDetail(fmt.Sprintf("Body size is exceeding the limit of '%d' bytes", maximumBytes))

		var actualRequestBody response.ProblemDetails
		requestBodyDecoder := json.NewDecoder(recorder.Body)
		err = requestBodyDecoder.Decode(&actualRequestBody)

		assert.NoError(t, err)

		assert.Equal(t, expectedResponseBody, actualRequestBody, fmt.Sprintf("Expected type to be '%v', but got '%v'", expectedResponseBody, actualRequestBody))
	})

	t.Run("Content smaller than limit", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		maximumBytes := int64(1)
		handler := func(w http.ResponseWriter, r *http.Request) {}
		middlewareHandler := middleware.RequestBodyLimit(logger, maximumBytes, handler)

		request, err := http.NewRequest("POST", "/", strings.NewReader(""))

		assert.NoError(t, err)

		recorder := httptest.NewRecorder()

		middlewareHandler.ServeHTTP(recorder, request)

		assert.NotEqual(t, http.StatusRequestEntityTooLarge, recorder.Code, fmt.Sprintf("Expected status code not to be '%d', but got '%d'", http.StatusRequestEntityTooLarge, recorder.Code))
	})
}
