package donetodos_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/query/donetodos"
	"github.com/matthiashermsen/dewit/domain"
)

func TestFromDomainTodos(t *testing.T) {
	expectedTodoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	expectedTodoTitle := "made-up"
	domainTodo, err := domain.NewTodo(expectedTodoID, expectedTodoTitle)

	assert.NoError(t, err)

	err = domainTodo.MarkAsDone()
	assert.NoError(t, err)

	responseBody := donetodos.FromDomainTodos([]*domain.Todo{domainTodo})

	assert.Len(t, responseBody, 1, "Expected response body to have exactly one item")

	actualTodo := responseBody[0]

	assert.Equal(t, expectedTodoID, actualTodo.ID, fmt.Sprintf("Expected todo ID to be '%s' but got '%s'", expectedTodoID, actualTodo.ID))
	assert.Equal(t, expectedTodoTitle, actualTodo.Title, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", expectedTodoTitle, actualTodo.Title))
}
