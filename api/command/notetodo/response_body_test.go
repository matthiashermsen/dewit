package notetodo_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/api/command/notetodo"
	"github.com/matthiashermsen/dewit/domain"
)

func TestFromDomainTodo(t *testing.T) {
	expectedTodoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	expectedTodoTitle := "made-up"
	domainTodo, err := domain.NewTodo(expectedTodoID, expectedTodoTitle)

	assert.NoError(t, err)

	responseBody := notetodo.FromDomainTodo(domainTodo)

	assert.Equal(t, expectedTodoID, responseBody.ID, fmt.Sprintf("Expected todo ID to be '%s' but got '%s'", expectedTodoID, responseBody.ID))
	assert.Equal(t, expectedTodoTitle, responseBody.Title, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", expectedTodoTitle, responseBody.Title))
}
