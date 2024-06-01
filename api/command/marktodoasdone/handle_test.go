package marktodoasdone_test

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

	"github.com/matthiashermsen/dewit/api/command/marktodoasdone"
	"github.com/matthiashermsen/dewit/api/response"
	"github.com/matthiashermsen/dewit/store"
)

func TestHandle(t *testing.T) {
	t.Run("Responds with 400 if request is invalid.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		handler := marktodoasdone.Handle(nil, logger)
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

	t.Run("Responds with 404 if todo was not found.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		mockStore := NotFoundMockStore{}

		handler := marktodoasdone.Handle(mockStore, logger)
		recorder := httptest.NewRecorder()

		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		requestBody := marktodoasdone.RequestBody{
			TodoID: todoID,
		}
		encodedRequestBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		request, err := http.NewRequest("POST", "", bytes.NewReader(encodedRequestBody))
		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusNotFound, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)
		assert.NoError(t, err)

		expectedProblemDetails := response.NewProblemDetails("TODO_NOT_FOUND", http.StatusNotFound, "Todo not found", fmt.Sprintf("Could not find todo with ID '%s'.", requestBody.TodoID))

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response to be '%v', but got '%v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Responds with 409 if todo is already marked as done.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		mockStore := TodoAlreadyMarkedAsDoneMockStore{}

		handler := marktodoasdone.Handle(mockStore, logger)
		recorder := httptest.NewRecorder()

		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		requestBody := marktodoasdone.RequestBody{
			TodoID: todoID,
		}
		encodedRequestBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		request, err := http.NewRequest("POST", "", bytes.NewReader(encodedRequestBody))
		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusConflict, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusConflict, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)
		assert.NoError(t, err)

		expectedProblemDetails := response.NewProblemDetails("TODO_ALREADY_MARKED_AS_DONE", http.StatusConflict, "Todo already marked as done", fmt.Sprintf("Todo with ID '%s' is already marked as done.", requestBody.TodoID))

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response to be '%v', but got '%v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Responds with internal error if something failed.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		mockStore := FailureMockStore{}

		handler := marktodoasdone.Handle(mockStore, logger)
		recorder := httptest.NewRecorder()

		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		requestBody := marktodoasdone.RequestBody{
			TodoID: todoID,
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

	t.Run("Marks todo as done.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		mockStore := MarkTodoAsDoneMockStore{}

		handler := marktodoasdone.Handle(mockStore, logger)
		recorder := httptest.NewRecorder()

		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		requestBody := marktodoasdone.RequestBody{
			TodoID: todoID,
		}
		encodedRequestBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		request, err := http.NewRequest("POST", "", bytes.NewReader(encodedRequestBody))
		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusCreated, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusCreated, recorder.Code))
		assert.Empty(t, recorder.Body.Bytes(), "Expected response body to be empty.")
	})
}

type NotFoundMockStore struct{}

func (s NotFoundMockStore) GetTodoByID(todoID uuid.UUID) (*store.Todo, error) {
	return nil, store.NewTodoNotFoundError(todoID)
}

func (s NotFoundMockStore) MarkTodoAsDone(todoID uuid.UUID) error {
	return nil
}

type TodoAlreadyMarkedAsDoneMockStore struct{}

func (s TodoAlreadyMarkedAsDoneMockStore) GetTodoByID(todoID uuid.UUID) (*store.Todo, error) {
	todo := store.NewTodo(todoID, "made-up", true)

	return todo, nil
}

func (s TodoAlreadyMarkedAsDoneMockStore) MarkTodoAsDone(todoID uuid.UUID) error {
	return nil
}

type FailureMockStore struct{}

func (s FailureMockStore) GetTodoByID(todoID uuid.UUID) (*store.Todo, error) {
	return nil, errors.New("something failed")
}

func (s FailureMockStore) MarkTodoAsDone(todoID uuid.UUID) error {
	return nil
}

type MarkTodoAsDoneMockStore struct{}

func (s MarkTodoAsDoneMockStore) GetTodoByID(todoID uuid.UUID) (*store.Todo, error) {
	todo := store.NewTodo(todoID, "made-up", false)

	return todo, nil
}

func (s MarkTodoAsDoneMockStore) MarkTodoAsDone(todoID uuid.UUID) error {
	return nil
}
