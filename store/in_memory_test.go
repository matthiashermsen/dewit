package store_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/store"
)

func TestNewInMemoryStore(t *testing.T) {
	store := store.NewInMemoryStore(slog.New(slog.Default().Handler()))

	assert.NotNil(t, store, "Expected store not to be nil")
}

func TestGetRemainingTodos(t *testing.T) {
	t.Run("No todos", func(t *testing.T) {
		store := store.NewInMemoryStore(slog.New(slog.Default().Handler()))
		todos, err := store.GetRemainingTodos()

		assert.NoError(t, err)
		assert.Empty(t, todos, fmt.Sprintf("Expected todos to be empty but got '%v'", todos))
	})

	t.Run("Remaining and done todos", func(t *testing.T) {
		todoTitle := "made-up"

		firstRemainingTodoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		firstRemainingTodo := store.NewTodo(firstRemainingTodoID, todoTitle, false)

		firstDoneTodoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		firstDoneTodo := store.NewTodo(firstDoneTodoID, todoTitle, true)

		secondRemainingTodoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		secondRemainingTodo := store.NewTodo(secondRemainingTodoID, todoTitle, false)

		secondDoneTodoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		secondDoneTodo := store.NewTodo(secondDoneTodoID, todoTitle, true)

		store := store.NewInMemoryStore(slog.New(slog.Default().Handler()))

		err = store.CreateTodo(firstRemainingTodo)
		assert.NoError(t, err)

		err = store.CreateTodo(firstDoneTodo)
		assert.NoError(t, err)

		err = store.CreateTodo(secondRemainingTodo)
		assert.NoError(t, err)

		err = store.CreateTodo(secondDoneTodo)
		assert.NoError(t, err)

		todos, err := store.GetRemainingTodos()
		assert.NoError(t, err)
		assert.Len(t, todos, 2, fmt.Sprintf("Expected todos to have exactly one element but got '%d'", len(todos)))

		firstTodoFromSlice := todos[0]
		assert.Equal(t, firstRemainingTodo, firstTodoFromSlice, fmt.Sprintf("Expected todo to be '%v' but got '%v'", firstRemainingTodo, firstTodoFromSlice))

		secondTodoFromSlice := todos[1]
		assert.Equal(t, secondRemainingTodo, secondTodoFromSlice, fmt.Sprintf("Expected todo to be '%v' but got '%v'", secondRemainingTodo, secondTodoFromSlice))
	})
}

func TestGetDoneTodos(t *testing.T) {
	t.Run("No todos", func(t *testing.T) {
		store := store.NewInMemoryStore(slog.New(slog.Default().Handler()))
		todos, err := store.GetDoneTodos()

		assert.NoError(t, err)
		assert.Empty(t, todos, fmt.Sprintf("Expected todos to be empty but got '%v'", todos))
	})

	t.Run("Done and remaining todos", func(t *testing.T) {
		todoTitle := "made-up"

		firstRemainingTodoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		firstRemainingTodo := store.NewTodo(firstRemainingTodoID, todoTitle, false)

		firstDoneTodoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		firstDoneTodo := store.NewTodo(firstDoneTodoID, todoTitle, true)

		secondRemainingTodoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		secondRemainingTodo := store.NewTodo(secondRemainingTodoID, todoTitle, false)

		secondDoneTodoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		secondDoneTodo := store.NewTodo(secondDoneTodoID, todoTitle, true)

		store := store.NewInMemoryStore(slog.New(slog.Default().Handler()))

		err = store.CreateTodo(firstRemainingTodo)
		assert.NoError(t, err)

		err = store.CreateTodo(firstDoneTodo)
		assert.NoError(t, err)

		err = store.CreateTodo(secondRemainingTodo)
		assert.NoError(t, err)

		err = store.CreateTodo(secondDoneTodo)
		assert.NoError(t, err)

		todos, err := store.GetDoneTodos()
		assert.NoError(t, err)
		assert.Len(t, todos, 2, fmt.Sprintf("Expected todos to have exactly one element but got '%d'", len(todos)))

		firstTodoFromSlice := todos[0]
		assert.Equal(t, firstDoneTodo, firstTodoFromSlice, fmt.Sprintf("Expected todo to be '%v' but got '%v'", firstDoneTodo, firstTodoFromSlice))

		secondTodoFromSlice := todos[1]
		assert.Equal(t, secondDoneTodo, secondTodoFromSlice, fmt.Sprintf("Expected todo to be '%v' but got '%v'", secondDoneTodo, secondTodoFromSlice))
	})
}

func TestCreateTodo(t *testing.T) {
	t.Run("Duplicate ID", func(t *testing.T) {
		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		newTodo := store.NewTodo(todoID, "made-up-1", false)
		todoStore := store.NewInMemoryStore(slog.New(slog.Default().Handler()))
		err = todoStore.CreateTodo(newTodo)
		assert.NoError(t, err)

		err = todoStore.CreateTodo(newTodo)
		assert.Error(t, err)

		var duplicateTodoIDError *store.DuplicateTodoIDError
		assert.ErrorAs(t, err, &duplicateTodoIDError, "Expected error to be DuplicateTodoIDError")
	})

	t.Run("Adds todo to store", func(t *testing.T) {
		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		newTodo := store.NewTodo(todoID, "made-up-1", false)
		todoStore := store.NewInMemoryStore(slog.New(slog.Default().Handler()))
		err = todoStore.CreateTodo(newTodo)
		assert.NoError(t, err)

		todos, err := todoStore.GetRemainingTodos()
		assert.NoError(t, err)
		assert.Len(t, todos, 1)

		todoFromSlice := todos[0]
		assert.Equal(t, newTodo, todoFromSlice, fmt.Sprintf("Expected todo to be '%v' but got '%v'", newTodo, todoFromSlice))
	})
}

func TestMarkTodoAsDone(t *testing.T) {
	t.Run("Todo not found", func(t *testing.T) {
		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		todoStore := store.NewInMemoryStore(slog.New(slog.Default().Handler()))
		err = todoStore.MarkTodoAsDone(todoID)
		assert.Error(t, err)

		var todoNotFoundError *store.TodoNotFoundError
		assert.ErrorAs(t, err, &todoNotFoundError, "Expected error to be TodoNotFoundError")
	})

	t.Run("Marks todo as done", func(t *testing.T) {
		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		newTodo := store.NewTodo(todoID, "made-up-1", false)
		todoStore := store.NewInMemoryStore(slog.New(slog.Default().Handler()))

		err = todoStore.CreateTodo(newTodo)
		assert.NoError(t, err)

		todoStore.MarkTodoAsDone(newTodo.GetID())
		assert.NoError(t, err)

		todos, err := todoStore.GetDoneTodos()
		assert.NoError(t, err)
		assert.Len(t, todos, 1)

		expectedTodoID := newTodo.GetID()

		todoFromSlice := todos[0]
		actualTodoID := todoFromSlice.GetID()

		assert.Equal(t, expectedTodoID, actualTodoID, fmt.Sprintf("Expected todo ID to be '%s' but got '%s'", expectedTodoID, actualTodoID))
		assert.True(t, todoFromSlice.IsMarkedAsDone(), "Expected todo to be marked as done")
	})
}
