package command

import (
	"github.com/google/uuid"

	"github.com/matthiashermsen/dewit/store"
)

type Store interface {
	GetTodoByID(todoID uuid.UUID) (*store.Todo, error)
	CreateTodo(todo *store.Todo) error
	MarkTodoAsDone(todoID uuid.UUID) error
}
