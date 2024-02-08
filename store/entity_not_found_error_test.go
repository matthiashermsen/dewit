package store_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/store"
)

func TestNewEntityNotFoundError(t *testing.T) {
	id := "made-up"

	entityNotFoundError := store.NewEntityNotFoundError(id)

	assert.NotNil(t, entityNotFoundError)
}

func TestEntityNotFoundError(t *testing.T) {
	id := "made-up"

	entityNotFoundError := store.NewEntityNotFoundError(id)
	actualErrMsg := entityNotFoundError.Error()
	expectedErrMsg := fmt.Sprintf("Entity with id '%s' not found.", id)

	assert.Equal(t, expectedErrMsg, actualErrMsg, fmt.Sprintf("Expected error message to be '%s' but got '%s'", expectedErrMsg, actualErrMsg))
}
