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

	"github.com/matthiashermsen/dewit/api/response"
	"github.com/matthiashermsen/dewit/api/responsetype"
	routenotetodopkg "github.com/matthiashermsen/dewit/api/route/command/notetodo"
	domainnotetodopkg "github.com/matthiashermsen/dewit/domain/command/notetodo"
	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

func TestHandle(t *testing.T) {
	t.Run("Responds with 400 if request is invalid.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		var noteTodo domainnotetodopkg.ServiceFunc = func(i domainnotetodopkg.Input) (*todo.Todo, error) {
			return nil, nil
		}
		handler := routenotetodopkg.Handle(logger, noteTodo)
		recorder := httptest.NewRecorder()

		invalidJSONBytes := []byte("{")
		request, err := http.NewRequest("POST", "/note-todo", bytes.NewReader(invalidJSONBytes))

		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusBadRequest, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		expectedProblemDetails := response.GetJSONParseErrorProblemDetails()

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Responds with 400 if todo title is empty.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		var noteTodo domainnotetodopkg.ServiceFunc = func(i domainnotetodopkg.Input) (*todo.Todo, error) {
			return nil, &todo.TodoTitleEmptyError{}
		}
		handler := routenotetodopkg.Handle(logger, noteTodo)
		recorder := httptest.NewRecorder()

		todoTitle := "made-up"

		requestBody := routenotetodopkg.RequestBody{TodoTitle: todoTitle}
		jsonPayload, err := json.Marshal(requestBody)
		request, err := http.NewRequest("POST", "/note-todo", bytes.NewReader(jsonPayload))

		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusBadRequest, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		expectedProblemDetails := response.
			NewProblemDetails().
			WithType("TODO_TITLE_EMPTY").
			WithStatus(http.StatusBadRequest).
			WithTitle("Todo title is empty").
			WithDetail("Todo title must not be empty.")

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Responds with internal server error if something failed.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		var noteTodo domainnotetodopkg.ServiceFunc = func(i domainnotetodopkg.Input) (*todo.Todo, error) {
			return nil, errors.New("something failed")
		}
		handler := routenotetodopkg.Handle(logger, noteTodo)
		recorder := httptest.NewRecorder()

		todoTitle := "made-up"

		requestBody := routenotetodopkg.RequestBody{TodoTitle: todoTitle}
		jsonPayload, err := json.Marshal(requestBody)
		request, err := http.NewRequest("POST", "/note-todo", bytes.NewReader(jsonPayload))

		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusInternalServerError, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		expectedProblemDetails := response.GetInternalServerErrorProblemDetails()

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Responds with status code 200 and todo.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		todoTitle := "made-up"
		var noteTodo domainnotetodopkg.ServiceFunc = func(i domainnotetodopkg.Input) (*todo.Todo, error) {
			todoID, err := uuid.NewUUID()

			assert.NoError(t, err)

			newTodo, err := todo.NewTodo(todoID, todoTitle)

			assert.NoError(t, err)

			return newTodo, err
		}
		handler := routenotetodopkg.Handle(logger, noteTodo)
		recorder := httptest.NewRecorder()

		requestBody := routenotetodopkg.RequestBody{TodoTitle: todoTitle}
		jsonPayload, err := json.Marshal(requestBody)
		request, err := http.NewRequest("POST", "/note-todo", bytes.NewReader(jsonPayload))

		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusOK, recorder.Code))

		var actualResponseBody responsetype.Todo
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		assert.Equal(t, todoTitle, actualResponseBody.Title, fmt.Sprintf("Expected todo title to be '%s', but got '%s'", todoTitle, actualResponseBody.Title))
	})
}
