package notetodo_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/service/command/notetodo"
	"github.com/matthiashermsen/dewit/store"
)

func TestHandle(t *testing.T) {
	t.Run("Notes todo", func(t *testing.T) {
		todoTitle := "made-up"
		input := notetodo.NewInput(todoTitle)
		store := MockStore{}

		todo, err := notetodo.Handle(input, store)
		assert.NoError(t, err)

		actualTodoTitle := todo.GetTitle()
		assert.Equal(t, todoTitle, actualTodoTitle, fmt.Sprintf("Expected todo title to be '%s' but got '%s'", todoTitle, actualTodoTitle))
		assert.False(t, todo.IsMarkedAsDone(), "Expected todo not to be marked as done")
	})
}

type MockStore struct{}

func (s MockStore) CreateTodo(todo *store.Todo) error {
	return nil
}
