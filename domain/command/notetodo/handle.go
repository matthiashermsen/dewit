package notetodo

import (
	"log/slog"

	"github.com/google/uuid"

	domaintodo "github.com/matthiashermsen/dewit/domain/entity/todo"
	storetodo "github.com/matthiashermsen/dewit/store/todo"
)

func Handle(logger *slog.Logger, todoStore TodoStore) ServiceFunc {
	return func(i Input) (*domaintodo.Todo, error) {
		todoID, err := uuid.NewUUID()

		if err != nil {
			return nil, err
		}

		newTodo, err := domaintodo.NewTodo(todoID, i.TodoTitle)

		if err != nil {
			return nil, err
		}

		storeTodo := storetodo.FromDomainTodo(newTodo)
		err = todoStore.SaveTodo(storeTodo)

		if err != nil {
			return nil, err
		}

		return newTodo, err
	}
}
