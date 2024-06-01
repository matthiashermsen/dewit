package domain_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/domain"
)

func TestNewTodoAlreadyMarkedAsDoneError(t *testing.T) {
	todoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	err = domain.NewTodoAlreadyMarkedAsDoneError(todoID)

	assert.NotNil(t, err)
}

func TestTodoAlreadyMarkedAsDoneError(t *testing.T) {
	todoID, err := uuid.NewUUID()

	assert.NoError(t, err)

	err = domain.NewTodoAlreadyMarkedAsDoneError(todoID)
	actualErrMsg := err.Error()
	expectedErrMsg := fmt.Sprintf("Todo with ID '%s' is already marked as done.", todoID)

	assert.Equal(t, expectedErrMsg, actualErrMsg, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrMsg, actualErrMsg))
}
