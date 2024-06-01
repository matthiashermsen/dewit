package donetodos_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/service/query/donetodos"
	"github.com/matthiashermsen/dewit/store"
)

func TestHandle(t *testing.T) {
	t.Run("Returns done todos", func(t *testing.T) {
		mockStore := MockStore{}

		todos, err := donetodos.Handle(mockStore)
		assert.NoError(t, err)
		assert.Len(t, todos, 1, "Expected todos to have exactly one todo")

		todoFromSlice := todos[0]

		assert.True(t, todoFromSlice.IsMarkedAsDone(), "Expected todo to be marked as done")
	})
}

type MockStore struct{}

func (s MockStore) GetDoneTodos() ([]*store.Todo, error) {
	todoID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	todo := store.NewTodo(todoID, "made-up", true)
	todos := []*store.Todo{todo}

	return todos, nil
}
