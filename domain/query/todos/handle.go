package todos

import (
	"log/slog"

	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

func Handle(logger *slog.Logger, todoStore TodoStore) ServiceFunc {
	return func() ([]*todo.Todo, error) {
		storeTodos, err := todoStore.GetTodos()

		if err != nil {
			return nil, err
		}

		domainTodos := make([]*todo.Todo, 0)

		for _, storeTodo := range storeTodos {
			domainTodo, err := storeTodo.ToDomainTodo()

			if err != nil {
				return nil, err
			}

			domainTodos = append(domainTodos, domainTodo)
		}

		return domainTodos, nil
	}
}
