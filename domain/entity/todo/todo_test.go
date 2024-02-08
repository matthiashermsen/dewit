package todo_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

func TestNewTodo(t *testing.T) {
	t.Run("Returns TodoTitleEmptyError if title is empty", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		newTodo, err := todo.NewTodo(todoID, "")

		assert.Nil(t, newTodo)

		_, isTodoTitleEmptyError := err.(*todo.TodoTitleEmptyError)
		assert.True(t, isTodoTitleEmptyError, "Expected error to be TodoTitleEmptyError")
	})

	t.Run("Returns Todo", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		todoTitle := "made-up"

		newTodo, err := todo.NewTodo(todoID, todoTitle)

		assert.Nil(t, err)
		assert.Equal(t, todoID, newTodo.ID)
		assert.Equal(t, todoTitle, newTodo.Title)
	})
}
