package marktodoasdone_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/service/command/marktodoasdone"
)

func TestNewInput(t *testing.T) {
	expectedTodoID, err := uuid.NewUUID()
	assert.NoError(t, err)

	input := marktodoasdone.NewInput(expectedTodoID)

	assert.Equal(t, expectedTodoID, input.TodoID, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", expectedTodoID, input.TodoID))
}
