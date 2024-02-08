package todos_test

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

	"github.com/matthiashermsen/dewit/api/response"
	"github.com/matthiashermsen/dewit/api/responsetype"
	routetodospkg "github.com/matthiashermsen/dewit/api/route/query/todos"
	"github.com/matthiashermsen/dewit/domain/entity/todo"
	domaintodospkg "github.com/matthiashermsen/dewit/domain/query/todos"
)

func TestHandle(t *testing.T) {
	t.Run("Responds with internal server error if something failed.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		var getTodos domaintodospkg.ServiceFunc = func() ([]*todo.Todo, error) {
			return nil, errors.New("Something failed.")
		}
		handler := routetodospkg.Handle(logger, getTodos)

		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/todos", nil)

		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusInternalServerError, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		expectedProblemDetails := response.GetInternalServerErrorProblemDetails()

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Responds with status code 200 and todos.", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		todoTitle := "made-up"
		domainTodo, err := todo.NewTodo(todoID, todoTitle)
		domainTodos := []*todo.Todo{domainTodo}

		var getTodos domaintodospkg.ServiceFunc = func() ([]*todo.Todo, error) {
			return domainTodos, nil
		}
		logger := slog.New(slog.Default().Handler())
		handler := routetodospkg.Handle(logger, getTodos)

		recorder := httptest.NewRecorder()
		request, err := http.NewRequest("GET", "/todos", nil)

		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusOK, recorder.Code))

		var actualResponseBody []responsetype.Todo
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		responseTodos := make([]responsetype.Todo, 0)

		for _, domainTodo := range domainTodos {
			responseTodo := responsetype.FromDomainTodo(domainTodo)
			responseTodos = append(responseTodos, responseTodo)
		}

		assert.Equal(t, responseTodos, actualResponseBody, fmt.Sprintf("Expected response to be '%#v', but got '%#v'", responseTodos, actualResponseBody))
	})
}
