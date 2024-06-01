package store_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/store"
)

func TestNewTodoNotFoundError(t *testing.T) {
	id, err := uuid.NewUUID()

	assert.NoError(t, err)

	todoNotFoundError := store.NewTodoNotFoundError(id)

	assert.NotNil(t, todoNotFoundError)
}

func TestEntityNotFoundError(t *testing.T) {
	id, err := uuid.NewUUID()

	assert.NoError(t, err)

	todoNotFoundError := store.NewTodoNotFoundError(id)
	actualErrMsg := todoNotFoundError.Error()
	expectedErrMsg := fmt.Sprintf("Todo with ID '%s' not found.", id)

	assert.Equal(t, expectedErrMsg, actualErrMsg, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrMsg, actualErrMsg))
}
