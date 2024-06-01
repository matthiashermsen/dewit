package remainingtodos_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/service/query/remainingtodos"
	"github.com/matthiashermsen/dewit/store"
)

func TestHandle(t *testing.T) {
	t.Run("Returns remaining todos", func(t *testing.T) {
		mockStore := MockStore{}

		todos, err := remainingtodos.Handle(mockStore)
		assert.NoError(t, err)
		assert.Len(t, todos, 1, "Expected todos to have exactly one todo")

		todoFromSlice := todos[0]

		assert.False(t, todoFromSlice.IsMarkedAsDone(), "Expected todo not to be marked as done")
	})
}

type MockStore struct{}

func (s MockStore) GetRemainingTodos() ([]*store.Todo, error) {
	todoID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	todo := store.NewTodo(todoID, "made-up", false)
	todos := []*store.Todo{todo}

	return todos, nil
}
