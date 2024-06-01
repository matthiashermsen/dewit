package notetodo

import (
	"github.com/google/uuid"
	"github.com/matthiashermsen/dewit/domain"
	"github.com/matthiashermsen/dewit/store"
)

func Handle(input Input, dataStore Store) (*domain.Todo, error) {
	todoID, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	todo, err := domain.NewTodo(todoID, input.TodoTitle)

	if err != nil {
		return nil, err
	}

	storeTodo := store.FromDomainTodo(todo)
	err = dataStore.CreateTodo(storeTodo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
