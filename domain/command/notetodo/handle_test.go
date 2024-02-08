package notetodo_test

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/domain/command/notetodo"
	"github.com/matthiashermsen/dewit/store/todo"
)

func TestHandle(t *testing.T) {
	t.Run("Returns any error.", func(t *testing.T) {
		todoStore := TodoStoreReturningAnyError{}
		noteTodo := notetodo.Handle(slog.New(slog.Default().Handler()), &todoStore)
		input := notetodo.Input{TodoTitle: "made-up"}

		notedTodo, err := noteTodo(input)

		assert.Error(t, err)
		assert.Nil(t, notedTodo)
	})

	t.Run("Notes todo.", func(t *testing.T) {
		todoStore := TodoStoreReturningNoError{}
		noteTodo := notetodo.Handle(slog.New(slog.Default().Handler()), &todoStore)
		todoTitle := "made-up"
		input := notetodo.Input{TodoTitle: todoTitle}

		notedTodo, err := noteTodo(input)

		assert.NoError(t, err)
		assert.NotNil(t, notedTodo)
		assert.Equal(t, todoTitle, notedTodo.Title)
	})
}

type TodoStoreReturningAnyError struct{}

func (s *TodoStoreReturningAnyError) SaveTodo(todo *todo.Todo) error {
	return errors.New("made-up")
}

type TodoStoreReturningNoError struct{}

func (s *TodoStoreReturningNoError) SaveTodo(todo *todo.Todo) error {
	return nil
}
