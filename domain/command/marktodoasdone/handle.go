package marktodoasdone

import (
	"log/slog"

	"github.com/matthiashermsen/dewit/domain/entity/todo"
	"github.com/matthiashermsen/dewit/store"
)

func Handle(logger *slog.Logger, todoStore TodoStore) ServiceFunc {
	return func(i Input) error {
		err := todoStore.DeleteTodoById(i.TodoID)

		if err == nil {
			return nil
		}

		_, isEntityNotFoundError := err.(*store.EntityNotFoundError)

		if isEntityNotFoundError {
			todoNotFoundError := todo.NewTodoNotFoundError(i.TodoID)

			return todoNotFoundError
		}

		return err
	}
}
