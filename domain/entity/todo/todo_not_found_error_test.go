package todo_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

func TestNewTodoNotFoundError(t *testing.T) {
	todoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	todoNotFoundError := todo.NewTodoNotFoundError(todoID)

	assert.NotNil(t, todoNotFoundError)
}

func TestTodoNotFoundError(t *testing.T) {
	todoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	todoNotFoundError := todo.NewTodoNotFoundError(todoID)
	actualErrMsg := todoNotFoundError.Error()
	expectedErrMsg := fmt.Sprintf("Todo with ID '%s' not found.", todoID)

	assert.Equal(t, expectedErrMsg, actualErrMsg, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrMsg, actualErrMsg))
}
