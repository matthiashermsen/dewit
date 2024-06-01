package store_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/domain"
	"github.com/matthiashermsen/dewit/store"
)

func TestNewTodo(t *testing.T) {
	todoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	todoTitle := "made-up"
	isTodoMarkedAsDone := true
	todo := store.NewTodo(todoID, todoTitle, isTodoMarkedAsDone)

	actualTodoID := todo.GetID()
	assert.Equal(t, todoID, actualTodoID, fmt.Sprintf("Expected todo ID to be '%s' but got '%s'", todoID, actualTodoID))

	actualTodoTitle := todo.GetTitle()
	assert.Equal(t, todoTitle, actualTodoTitle, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", todoTitle, actualTodoTitle))

	actualIsTodoMarkedAsDone := todo.IsMarkedAsDone()
	assert.Equal(t, isTodoMarkedAsDone, actualIsTodoMarkedAsDone, "Expected todo to be marked as done")
}

func TestGetID(t *testing.T) {
	expectedTodoID, err := uuid.NewUUID()
	assert.NoError(t, err)

	newTodo := store.NewTodo(expectedTodoID, "made-up", false)

	actualTodoID := newTodo.GetID()
	assert.Equal(t, expectedTodoID, actualTodoID, fmt.Sprintf("Expected todo ID to be '%s' but got '%s'", expectedTodoID, actualTodoID))
}

func TestGetTitle(t *testing.T) {
	todoID, err := uuid.NewUUID()
	assert.NoError(t, err)

	expectedTodoTitle := "made-up"

	newTodo := store.NewTodo(todoID, expectedTodoTitle, false)

	actualTodoTitle := newTodo.GetTitle()
	assert.Equal(t, expectedTodoTitle, actualTodoTitle, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", expectedTodoTitle, actualTodoTitle))
}

func TestIsMarkedAsDone(t *testing.T) {
	todoID, err := uuid.NewUUID()
	assert.NoError(t, err)

	newTodo := store.NewTodo(todoID, "made-up", true)

	isTodoMarkedAsDone := newTodo.IsMarkedAsDone()
	assert.Equal(t, true, isTodoMarkedAsDone, "Expected todo to be marked as done")
}

func TestMarkAsDone(t *testing.T) {
	todoID, err := uuid.NewUUID()
	assert.NoError(t, err)

	newTodo := store.NewTodo(todoID, "made-up", false)
	newTodo.MarkAsDone()

	isTodoMarkedAsDone := newTodo.IsMarkedAsDone()
	assert.Equal(t, true, isTodoMarkedAsDone, "Expected todo to be marked as done")
}

func TestFromDomainTodo(t *testing.T) {
	todoID, err := uuid.NewUUID()
	assert.NoError(t, err)

	todoTitle := "made-up"
	domainTodo, err := domain.NewTodo(todoID, todoTitle)
	assert.NoError(t, err)

	storeTodo := store.FromDomainTodo(domainTodo)

	domainTodoID := domainTodo.GetID()
	storeTodoID := storeTodo.GetID()
	assert.Equal(t, domainTodoID, storeTodoID, fmt.Sprintf("Expected todo ID to be '%s' but got '%s'", domainTodoID, storeTodoID))

	domainTodoTitle := domainTodo.GetTitle()
	storeTodoTitle := storeTodo.GetTitle()
	assert.Equal(t, domainTodoTitle, storeTodoTitle, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", domainTodoTitle, storeTodoTitle))

	isDomainTodoMarkedAsDone := domainTodo.IsMarkedAsDone()
	isStoreTodoMarkedAsDone := storeTodo.IsMarkedAsDone()
	assert.Equal(t, isDomainTodoMarkedAsDone, isStoreTodoMarkedAsDone, "Invalid state for 'isMarkedAsDone'")
}

func TestToDomainTodo(t *testing.T) {
	todoID, err := uuid.NewUUID()
	assert.NoError(t, err)

	todoTitle := "made-up"
	isMarkedAsDone := true
	storeTodo := store.NewTodo(todoID, todoTitle, isMarkedAsDone)
	domainTodo, err := storeTodo.ToDomainTodo()

	assert.NoError(t, err)

	storeTodoID := storeTodo.GetID()
	domainTodoID := domainTodo.GetID()
	assert.Equal(t, storeTodoID, domainTodoID, fmt.Sprintf("Expected todo ID to be '%s' but got '%s'", storeTodoID, domainTodoID))

	domainTodoTitle := domainTodo.GetTitle()
	storeTodoTitle := storeTodo.GetTitle()
	assert.Equal(t, storeTodoTitle, domainTodoTitle, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", storeTodoTitle, domainTodoTitle))

	isDomainTodoMarkedAsDone := domainTodo.IsMarkedAsDone()
	isStoreTodoMarkedAsDone := storeTodo.IsMarkedAsDone()
	assert.Equal(t, isStoreTodoMarkedAsDone, isDomainTodoMarkedAsDone, "Invalid state for 'isMarkedAsDone'")
}
