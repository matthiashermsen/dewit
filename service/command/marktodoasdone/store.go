package marktodoasdone

import (
	"github.com/google/uuid"

	"github.com/matthiashermsen/dewit/store"
)

type Store interface {
	GetTodoByID(todoID uuid.UUID) (*store.Todo, error)
	MarkTodoAsDone(todoID uuid.UUID) error
}
