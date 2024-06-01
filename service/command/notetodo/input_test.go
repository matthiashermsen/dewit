package notetodo_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/service/command/notetodo"
)

func TestNewInput(t *testing.T) {
	expectedTodoTitle := "made-up"

	input := notetodo.NewInput(expectedTodoTitle)

	assert.Equal(t, expectedTodoTitle, input.TodoTitle, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", expectedTodoTitle, input.TodoTitle))
}
