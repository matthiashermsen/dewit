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

	"github.com/matthiashermsen/dewit/api/response"
	routemarktodoasdonepkg "github.com/matthiashermsen/dewit/api/route/command/marktodoasdone"
	domainmarktodoasdonepkg "github.com/matthiashermsen/dewit/domain/command/marktodoasdone"
	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

func TestHandle(t *testing.T) {
	t.Run("Responds with 400 if request is invalid.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		var markTodoAsDone domainmarktodoasdonepkg.ServiceFunc = func(i domainmarktodoasdonepkg.Input) error {
			return nil
		}
		handler := routemarktodoasdonepkg.Handle(logger, markTodoAsDone)
		recorder := httptest.NewRecorder()

		invalidJSONBytes := []byte("{")
		request, err := http.NewRequest("POST", "/mark-todo-as-done", bytes.NewReader(invalidJSONBytes))

		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusBadRequest, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusBadRequest, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		expectedProblemDetails := response.GetJSONParseErrorProblemDetails()

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Responds with 404 if todo doesn't exist.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		var markTodoAsDone domainmarktodoasdonepkg.ServiceFunc = func(i domainmarktodoasdonepkg.Input) error {
			return todo.NewTodoNotFoundError(i.TodoID)
		}
		handler := routemarktodoasdonepkg.Handle(logger, markTodoAsDone)
		recorder := httptest.NewRecorder()
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		requestBody := routemarktodoasdonepkg.RequestBody{TodoID: todoID}
		jsonPayload, err := json.Marshal(requestBody)
		request, err := http.NewRequest("POST", "/mark-todo-as-done", bytes.NewReader(jsonPayload))

		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusNotFound, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusNotFound, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		expectedProblemDetails := response.
			NewProblemDetails().
			WithType("TODO_NOT_FOUND").
			WithStatus(http.StatusNotFound).
			WithTitle("Todo not found").
			WithDetail(fmt.Sprintf("Todo with ID '%s' wasn't found.", todoID))

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Responds with internal server error if something failed.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		var markTodoAsDone domainmarktodoasdonepkg.ServiceFunc = func(i domainmarktodoasdonepkg.Input) error {
			return errors.New("something failed")
		}
		handler := routemarktodoasdonepkg.Handle(logger, markTodoAsDone)
		recorder := httptest.NewRecorder()
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		requestBody := routemarktodoasdonepkg.RequestBody{TodoID: todoID}
		jsonPayload, err := json.Marshal(requestBody)
		request, err := http.NewRequest("POST", "/mark-todo-as-done", bytes.NewReader(jsonPayload))

		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusInternalServerError, recorder.Code))

		var actualResponseBody response.ProblemDetails
		err = json.Unmarshal(recorder.Body.Bytes(), &actualResponseBody)

		assert.NoError(t, err)

		expectedProblemDetails := response.GetInternalServerErrorProblemDetails()

		assert.Equal(t, expectedProblemDetails, actualResponseBody, fmt.Sprintf("Expected response details to be '%#v', but got '%#v'", expectedProblemDetails, actualResponseBody))
	})

	t.Run("Responds with status code 200.", func(t *testing.T) {
		logger := slog.New(slog.Default().Handler())
		var markTodoAsDone domainmarktodoasdonepkg.ServiceFunc = func(i domainmarktodoasdonepkg.Input) error {
			return nil
		}
		handler := routemarktodoasdonepkg.Handle(logger, markTodoAsDone)
		recorder := httptest.NewRecorder()
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		requestBody := routemarktodoasdonepkg.RequestBody{TodoID: todoID}
		jsonPayload, err := json.Marshal(requestBody)
		request, err := http.NewRequest("POST", "/mark-todo-as-done", bytes.NewReader(jsonPayload))

		assert.NoError(t, err)

		handler.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code, fmt.Sprintf("Expected status code to be '%d', but got '%d'", http.StatusOK, recorder.Code))
	})
}
