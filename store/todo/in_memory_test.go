package todo_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/store"
	"github.com/matthiashermsen/dewit/store/todo"
)

func TestNewInMemoryStore(t *testing.T) {
	store := todo.NewInMemoryStore(slog.New(slog.Default().Handler()))
	todos, err := store.GetTodos()

	assert.NoError(t, err)
	assert.Empty(t, todos, fmt.Sprintf("Expected store todos to be empty but got '%v'", todos))
}

func TestGetTodos(t *testing.T) {
	t.Run("No todos", func(t *testing.T) {
		store := todo.NewInMemoryStore(slog.New(slog.Default().Handler()))
		todos, err := store.GetTodos()

		assert.NoError(t, err)
		assert.Empty(t, todos, fmt.Sprintf("Expected todos to be empty but got '%v'", todos))
	})

	t.Run("One todo", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		todoTitle := "made-up"
		newTodo := todo.NewTodo(todoID, todoTitle)

		store := todo.NewInMemoryStore(slog.New(slog.Default().Handler()))
		err = store.SaveTodo(newTodo)

		assert.NoError(t, err)

		todos, err := store.GetTodos()

		assert.NoError(t, err)
		assert.Len(t, todos, 1, fmt.Sprintf("Expected todos to have exactly one element but got '%d'", len(todos)))

		todoFromSlice := todos[0]

		assert.Equal(t, newTodo, todoFromSlice, fmt.Sprintf("Expected todo to be '%v' but got '%v'", newTodo, todoFromSlice))
	})
}

func TestSaveTodo(t *testing.T) {
	t.Run("Duplicate ID", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		newTodo := todo.NewTodo(todoID, "made-up-1")
		todoStore := todo.NewInMemoryStore(slog.New(slog.Default().Handler()))
		err = todoStore.SaveTodo(newTodo)

		assert.NoError(t, err)

		err = todoStore.SaveTodo(newTodo)

		assert.Error(t, err)

		_, isDuplicateEntityIdError := err.(*store.DuplicateEntityIdError)
		assert.True(t, isDuplicateEntityIdError, "Expected error to be DuplicateEntityIdError")
	})

	t.Run("Adds to store", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		newTodo := todo.NewTodo(todoID, "made-up-1")
		store := todo.NewInMemoryStore(slog.New(slog.Default().Handler()))
		err = store.SaveTodo(newTodo)

		assert.NoError(t, err)

		todos, err := store.GetTodos()

		assert.NoError(t, err)
		assert.Len(t, todos, 1)

		todoFromSlice := todos[0]

		assert.Equal(t, newTodo, todoFromSlice, fmt.Sprintf("Expected todo to be '%v' but got '%v'", newTodo, todoFromSlice))
	})
}

func TestDeleteTodoById(t *testing.T) {
	t.Run("Not found", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		todoStore := todo.NewInMemoryStore(slog.New(slog.Default().Handler()))

		err = todoStore.DeleteTodoById(todoID)

		assert.Error(t, err)

		_, isEntityNotFoundError := err.(*store.EntityNotFoundError)
		assert.True(t, isEntityNotFoundError, "Expected error to be EntityNotFoundError")
	})

	t.Run("Deletes todo", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		newTodo := todo.NewTodo(todoID, "made-up-1")
		store := todo.NewInMemoryStore(slog.New(slog.Default().Handler()))
		err = store.SaveTodo(newTodo)

		assert.NoError(t, err)

		err = store.DeleteTodoById(todoID)

		assert.NoError(t, err)

		todos, err := store.GetTodos()

		assert.NoError(t, err)
		assert.Len(t, todos, 0)
	})
}
