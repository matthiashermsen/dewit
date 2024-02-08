package marktodoasdone_test

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/matthiashermsen/dewit/domain/command/marktodoasdone"
	"github.com/matthiashermsen/dewit/domain/entity/todo"
	"github.com/matthiashermsen/dewit/store"
)

func TestHandle(t *testing.T) {
	t.Run("Returns TodoNotFoundError on EntityNotFoundError.", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		todoStore := TodoStoreReturningEntityNotFoundError{}
		markTodoAsDone := marktodoasdone.Handle(slog.New(slog.Default().Handler()), &todoStore)

		input := marktodoasdone.Input{TodoID: todoID}

		err = markTodoAsDone(input)

		_, isTodoNotFoundError := err.(*todo.TodoNotFoundError)
		assert.True(t, isTodoNotFoundError, "Expected error to be TodoNotFoundError")
	})

	t.Run("Returns any error.", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		todoStore := TodoStoreReturningAnyError{}
		markTodoAsDone := marktodoasdone.Handle(slog.New(slog.Default().Handler()), &todoStore)

		input := marktodoasdone.Input{TodoID: todoID}

		err = markTodoAsDone(input)

		assert.Error(t, err)
	})

	t.Run("Marks todo as done.", func(t *testing.T) {
		todoID, err := uuid.NewUUID()

		assert.NoError(t, err)

		todoStore := TodoStoreReturningNoError{}
		markTodoAsDone := marktodoasdone.Handle(slog.New(slog.Default().Handler()), &todoStore)

		input := marktodoasdone.Input{TodoID: todoID}

		err = markTodoAsDone(input)

		assert.NoError(t, err)
	})
}

type TodoStoreReturningEntityNotFoundError struct{}

func (s *TodoStoreReturningEntityNotFoundError) DeleteTodoById(todoID uuid.UUID) error {
	return store.NewEntityNotFoundError(todoID.String())
}

type TodoStoreReturningAnyError struct{}

func (s *TodoStoreReturningAnyError) DeleteTodoById(todoID uuid.UUID) error {
	return errors.New("made-up")
}

type TodoStoreReturningNoError struct{}

func (s *TodoStoreReturningNoError) DeleteTodoById(todoID uuid.UUID) error {
	return nil
}
