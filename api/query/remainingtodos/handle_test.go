package remainingtodos_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/query/remainingtodos"
	"github.com/matthiashermsen/dewit/api/response"
	"github.com/matthiashermsen/dewit/store"
)

func TestHandle(t *testing.T) {
	t.Run("Responds with internal error if something failed.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		mockStore := FailureMockStore{}

		handler := remainingtodos.Handle(mockStore, logger)
		recorder := httptest.NewRecorder()

		request, err := http.NewRequest("GET", "", nil)
		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusInternalServerError, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)
		assert.NoError(t, err)

		expectedProblemDetails := response.NewProblemDetails("INTERNAL_ERROR", http.StatusInternalServerError, "Internal error", "The server encountered an unexpected condition that prevented it from fulfilling the request.")

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response to be '%v', but got '%v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Returns todos.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())

		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		todo := store.NewTodo(todoID, "made-up", false)
		todos := []*store.Todo{todo}

		mockStore := NewMockStore(todos)

		handler := remainingtodos.Handle(mockStore, logger)
		recorder := httptest.NewRecorder()

		request, err := http.NewRequest("GET", "", nil)
		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusOK, recorder.Code))

		var actualResponseBody remainingtodos.ResponseBody
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)
		assert.NoError(t, err)

		assert.Len(t, actualResponseBody, 1, "Expected response body to have exactly one todo")

		expectedTodo := todos[0]
		actualTodo := actualResponseBody[0]

		assert.Equal(t, expectedTodo.GetID(), actualTodo.ID, fmt.Sprintf("Expected todo ID to be '%s', but got '%s'", expectedTodo.GetID(), actualTodo.ID))
		assert.Equal(t, expectedTodo.GetTitle(), actualTodo.Title, fmt.Sprintf("Expected todo title to be '%s', but got '%s'", expectedTodo.GetTitle(), actualTodo.Title))
	})
}

type FailureMockStore struct{}

func (s FailureMockStore) GetRemainingTodos() ([]*store.Todo, error) {
	return nil, errors.New("something failed")
}

type MockStore struct {
	Todos []*store.Todo
}

func NewMockStore(todos []*store.Todo) *MockStore {
	return &MockStore{Todos: todos}
}

func (s *MockStore) GetRemainingTodos() ([]*store.Todo, error) {
	return s.Todos, nil
}
