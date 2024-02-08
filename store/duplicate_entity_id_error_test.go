package store_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/store"
)

func TestNewDuplicateEntityIdError(t *testing.T) {
	id := "made-up"

	duplicateEntityIdError := store.NewDuplicateEntityIdError(id)

	assert.NotNil(t, duplicateEntityIdError)
}

func TestDuplicateEntityIdError(t *testing.T) {
	id := "made-up"

	duplicateEntityIdError := store.NewDuplicateEntityIdError(id)
	actualErrMsg := duplicateEntityIdError.Error()
	expectedErrMsg := fmt.Sprintf("Duplicate ID '%s' for entity.", id)

	assert.Equal(t, expectedErrMsg, actualErrMsg, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrMsg, actualErrMsg))
}
