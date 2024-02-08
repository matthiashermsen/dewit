package todo_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	domaintodo "github.com/matthiashermsen/dewit/domain/entity/todo"
	storetodo "github.com/matthiashermsen/dewit/store/todo"
)

func TestNewTodo(t *testing.T) {
	todoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	todoTitle := "made-up"
	todo := storetodo.NewTodo(todoID, todoTitle)

	assert.Equal(t, todoID, todo.ID, fmt.Sprintf("Expected todo id to be '%s' but got '%s'", todoID, todo.ID))
	assert.Equal(t, todoTitle, todo.Title, fmt.Sprintf("Expected todo id to be '%s' but got '%s'", todoTitle, todo.Title))
}

func TestFromDomainTodo(t *testing.T) {
	todoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	todoTitle := "made-up"
	domainTodo, err := domaintodo.NewTodo(todoID, todoTitle)

	assert.NoError(t, err)

	storeTodo := storetodo.FromDomainTodo(domainTodo)

	assert.Equal(t, domainTodo.ID, storeTodo.ID, fmt.Sprintf("Expected todo ID to be '%v' but got '%v'", domainTodo.ID, storeTodo.ID))
	assert.Equal(t, domainTodo.Title, storeTodo.Title, fmt.Sprintf("Expected todo to be '%s' but got '%s'", domainTodo.Title, storeTodo.Title))
}

func TestToDomainTodo(t *testing.T) {
	todoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	todoTitle := "made-up"
	storeTodo := storetodo.NewTodo(todoID, todoTitle)
	domainTodo, err := storeTodo.ToDomainTodo()

	assert.NoError(t, err)
	assert.Equal(t, domainTodo.ID, storeTodo.ID, fmt.Sprintf("Expected todo ID to be '%v' but got '%v'", domainTodo.ID, storeTodo.ID))
	assert.Equal(t, domainTodo.Title, storeTodo.Title, fmt.Sprintf("Expected todo to be '%s' but got '%s'", domainTodo.Title, storeTodo.Title))
}
