package responsetype_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/responsetype"
	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

func TestFromDomainTodo(t *testing.T) {
	todoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	todoTitle := "made-up"
	domainTodo, err := todo.NewTodo(todoID, todoTitle)

	assert.NoError(t, err)

	responseTodo := responsetype.FromDomainTodo(domainTodo)

	assert.Equal(t, domainTodo.ID, responseTodo.ID, fmt.Sprintf("Expected todo ID to be '%v' but got '%v'", domainTodo.ID, responseTodo.ID))
	assert.Equal(t, domainTodo.Title, responseTodo.Title, fmt.Sprintf("Expected todo title to be '%v' but got '%v'", domainTodo.Title, responseTodo.Title))
}
