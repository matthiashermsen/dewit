package todo_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

func TestTodoTitleEmptyError(t *testing.T) {
	todoTitleEmptyError := todo.TodoTitleEmptyError{}
	actualErrMsg := todoTitleEmptyError.Error()
	expectedErrMsg := "Todo title must not be empty."

	assert.Equal(t, expectedErrMsg, actualErrMsg, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrMsg, actualErrMsg))
}
