package notetodo_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/command/notetodo"
	"github.com/matthiashermsen/dewit/api/response"
	"github.com/matthiashermsen/dewit/domain"
	"github.com/matthiashermsen/dewit/store"
)

func TestHandle(t *testing.T) {
	t.Run("Responds with 400 if request is invalid.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		handler := notetodo.Handle(nil, logger)
		recorder := httptest.NewRecorder()

		invalidJSONBytes := []byte("{")

		request, err := http.NewRequest("POST", "", bytes.NewReader(invalidJSONBytes))
		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusBadRequest, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)
		assert.NoError(t, err)

		expectedProblemDetails := response.NewProblemDetails("REQUEST_BODY_INVALID", http.StatusBadRequest, "Request body invalid", "")

		assert.Equal(t, expectedProblemDetails.Type, actualResponseBody.Type, fmt.Sprintf("Expected type to be '%s', but got '%s'", expectedProblemDetails.Type, actualResponseBody.Type))
		assert.Equal(t, expectedProblemDetails.Status, actualResponseBody.Status, fmt.Sprintf("Expected status to be '%d', but got '%d'", expectedProblemDetails.Status, actualResponseBody.Status))
		assert.Equal(t, expectedProblemDetails.Title, actualResponseBody.Title, fmt.Sprintf("Expected title to be '%s', but got '%s'", expectedProblemDetails.Title, actualResponseBody.Title))
	})

	t.Run("Responds with 400 if todo title is empty.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		handler := notetodo.Handle(nil, logger)
		recorder := httptest.NewRecorder()

		requestBody := notetodo.RequestBody{
			TodoTitle: "",
		}
		encodedRequestBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		request, err := http.NewRequest("POST", "", bytes.NewReader(encodedRequestBody))
		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusBadRequest, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)
		assert.NoError(t, err)

		expectedProblemDetails := response.NewProblemDetails("TODO_TITLE_EMPTY", http.StatusBadRequest, "Todo title is empty", "Todo title must not be empty.")

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response to be '%v', but got '%v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Responds with internal error if something failed.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		mockStore := FailureMockStore{}

		handler := notetodo.Handle(mockStore, logger)
		recorder := httptest.NewRecorder()

		requestBody := notetodo.RequestBody{
			TodoTitle: "made-up",
		}
		encodedRequestBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		request, err := http.NewRequest("POST", "", bytes.NewReader(encodedRequestBody))
		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusInternalServerError, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)
		assert.NoError(t, err)

		expectedProblemDetails := response.NewProblemDetails("INTERNAL_ERROR", http.StatusInternalServerError, "Internal error", "The server encountered an unexpected condition that prevented it from fulfilling the request.")

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response to be '%v', but got '%v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Notes todo.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		mockStore := MockStore{}

		handler := notetodo.Handle(mockStore, logger)
		recorder := httptest.NewRecorder()

		requestBody := notetodo.RequestBody{
			TodoTitle: "made-up",
		}
		encodedRequestBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		request, err := http.NewRequest("POST", "", bytes.NewReader(encodedRequestBody))
		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusCreated, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusCreated, recorder.Code))

		todo, err := domain.NewTodo(uuid.Nil, requestBody.TodoTitle)
		assert.NoError(t, err)

		var actualResponseBody notetodo.ResponseBody
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)
		assert.NoError(t, err)

		assert.NotNil(t, actualResponseBody.ID, "Expected todo ID not to be nil")
		assert.Equal(t, todo.GetTitle(), actualResponseBody.Title, fmt.Sprintf("Expected todo title to be '%s', but got '%s'", todo.GetTitle(), actualResponseBody.Title))
	})
}

type MockStore struct{}

func (s MockStore) CreateTodo(todo *store.Todo) error {
	return nil
}

type FailureMockStore struct{}

func (s FailureMockStore) CreateTodo(todo *store.Todo) error {
	return errors.New("Something failed")
}
