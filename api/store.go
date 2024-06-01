package api

import (
	"github.com/google/uuid"

	"github.com/matthiashermsen/dewit/store"
)

type Store interface {
	GetRemainingTodos() ([]*store.Todo, error)
	GetDoneTodos() ([]*store.Todo, error)
	GetTodoByID(todoID uuid.UUID) (*store.Todo, error)
	CreateTodo(todo *store.Todo) error
	MarkTodoAsDone(todoID uuid.UUID) error
}
