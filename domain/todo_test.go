package domain_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/domain"
)

func TestNewTodo(t *testing.T) {
	t.Run("Returns TodoTitleEmptyError if title is empty", func(t *testing.T) {
		expectedTodoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		newTodo, err := domain.NewTodo(expectedTodoID, "")

		assert.Nil(t, newTodo, "Expected todo to be nil")
		assert.ErrorIs(t, err, domain.ErrTodoTitleEmpty, "Expected error to be ErrTodoTitleEmpty")
	})

	t.Run("Returns Todo", func(t *testing.T) {
		expectedTodoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		expectedTodoTitle := "made-up"
		newTodo, err := domain.NewTodo(expectedTodoID, expectedTodoTitle)

		assert.NoError(t, err)

		actualTodoID := newTodo.GetID()
		assert.Equal(t, expectedTodoID, actualTodoID, fmt.Sprintf("Expected todo ID to be '%s' but got '%s'", expectedTodoID, actualTodoID))

		actualTodoTitle := newTodo.GetTitle()
		assert.Equal(t, expectedTodoTitle, actualTodoTitle, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", expectedTodoTitle, actualTodoTitle))

		assert.Equal(t, false, newTodo.IsMarkedAsDone(), "Expected todo not to be marked as done")
	})
}

func TestGetID(t *testing.T) {
	expectedTodoID, err := uuid.NewUUID()
	assert.NoError(t, err)

	newTodo, err := domain.NewTodo(expectedTodoID, "made-up")

	assert.NoError(t, err)

	actualTodoID := newTodo.GetID()
	assert.Equal(t, expectedTodoID, actualTodoID, fmt.Sprintf("Expected todo ID to be '%s' but got '%s'", expectedTodoID, actualTodoID))
}

func TestGetTitle(t *testing.T) {
	todoID, err := uuid.NewUUID()
	assert.NoError(t, err)

	expectedTodoTitle := "made-up"

	newTodo, err := domain.NewTodo(todoID, expectedTodoTitle)

	assert.NoError(t, err)

	actualTodoTitle := newTodo.GetTitle()
	assert.Equal(t, expectedTodoTitle, actualTodoTitle, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", expectedTodoTitle, actualTodoTitle))
}

func TestIsMarkedAsDone(t *testing.T) {
	todoID, err := uuid.NewUUID()
	assert.NoError(t, err)

	newTodo, err := domain.NewTodo(todoID, "made-up")
	assert.NoError(t, err)

	err = newTodo.MarkAsDone()
	assert.NoError(t, err)

	isTodoMarkedAsDone := newTodo.IsMarkedAsDone()
	assert.Equal(t, true, isTodoMarkedAsDone, "Expected todo to be marked as done")
}

func TestMarkAsDone(t *testing.T) {
	t.Run("Returns TodoAlreadyMarkedAsDoneError if todo is already marked as done", func(t *testing.T) {
		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		newTodo, err := domain.NewTodo(todoID, "made-up")
		assert.NoError(t, err)

		err = newTodo.MarkAsDone()
		assert.NoError(t, err)

		err = newTodo.MarkAsDone()

		var todoAlreadyMarkedAsDoneError *domain.TodoAlreadyMarkedAsDoneError
		assert.ErrorAs(t, err, &todoAlreadyMarkedAsDoneError, "Expected error to be TodoAlreadyMarkedAsDoneError")
	})

	t.Run("Marks todo as done", func(t *testing.T) {
		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		newTodo, err := domain.NewTodo(todoID, "made-up")
		assert.NoError(t, err)

		err = newTodo.MarkAsDone()

		assert.NoError(t, err)
		assert.Equal(t, true, newTodo.IsMarkedAsDone(), "Expected todo to be marked as done")
	})
}
