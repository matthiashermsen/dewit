package marktodoasdone

import "github.com/google/uuid"

type TodoStore interface {
	DeleteTodoById(todoID uuid.UUID) error
}
