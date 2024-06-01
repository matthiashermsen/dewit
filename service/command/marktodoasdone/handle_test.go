package marktodoasdone_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/service/command/marktodoasdone"
	"github.com/matthiashermsen/dewit/store"
)

func TestHandle(t *testing.T) {
	t.Run("Marks todo as done", func(t *testing.T) {
		todoID, err := uuid.NewUUID()
		assert.NoError(t, err)

		input := marktodoasdone.NewInput(todoID)
		mockStore := MockStore{}

		err = marktodoasdone.Handle(input, mockStore)
		assert.NoError(t, err)
	})
}

type MockStore struct{}

func (s MockStore) GetTodoByID(todoID uuid.UUID) (*store.Todo, error) {
	todo := store.NewTodo(todoID, "made-up", false)

	return todo, nil
}

func (s MockStore) MarkTodoAsDone(todoID uuid.UUID) error {
	return nil
}
