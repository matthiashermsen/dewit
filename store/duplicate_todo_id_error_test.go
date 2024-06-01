package store_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/store"
)

func TestNewDuplicateTodoIDError(t *testing.T) {
	id, err := uuid.NewUUID()

	assert.NoError(t, err)

	duplicateTodoIDError := store.NewDuplicateTodoIDError(id)

	assert.NotNil(t, duplicateTodoIDError)
}

func TestDuplicateTodoIDError(t *testing.T) {
	id, err := uuid.NewUUID()

	assert.NoError(t, err)

	duplicateTodoIDError := store.NewDuplicateTodoIDError(id)
	actualErrMsg := duplicateTodoIDError.Error()
	expectedErrMsg := fmt.Sprintf("Duplicate todo ID '%s'.", id)

	assert.Equal(t, expectedErrMsg, actualErrMsg, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrMsg, actualErrMsg))
}
