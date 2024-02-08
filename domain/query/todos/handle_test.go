package todos_test

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/domain/query/todos"
	"github.com/matthiashermsen/dewit/store/todo"
)

func TestHandle(t *testing.T) {
	t.Run("Returns nil for todos and error if something fails.", func(t *testing.T) {
		todoStore := TodoStoreReturningError{}
		queryTodos := todos.Handle(slog.New(slog.Default().Handler()), &todoStore)

		todos, err := queryTodos()

		assert.Nil(t, todos, "Expected todos to be nil")
		assert.Error(t, err, "Expected error not to be nil")
	})

	t.Run("Returns todos.", func(t *testing.T) {
		todoStore := TodoStoreReturningTodos{}
		queryTodos := todos.Handle(slog.New(slog.Default().Handler()), &todoStore)

		todos, err := queryTodos()

		assert.Nil(t, err, "Expected error to be nil")
		assert.NotNil(t, todos, "Expected todos not to be nil")
		assert.Len(t, todos, 1, "Expected todos to have exactly one element")

		retrievedTodo := todos[0]

		assert.Equal(t, "made-up", retrievedTodo.Title)
	})
}

type TodoStoreReturningError struct{}

func (store *TodoStoreReturningError) GetTodos() ([]*todo.Todo, error) {
	return nil, errors.New("Something failed")
}

type TodoStoreReturningTodos struct{}

func (store *TodoStoreReturningTodos) GetTodos() ([]*todo.Todo, error) {
	todoID, err := uuid.NewUUID()
	todoTitle := "made-up"

	if err != nil {
		return nil, err
	}

	newTodo := todo.NewTodo(todoID, todoTitle)
	todos := []*todo.Todo{newTodo}

	return todos, nil
}
