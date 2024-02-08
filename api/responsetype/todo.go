package responsetype

import (
	"github.com/google/uuid"

	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

type Todo struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

func FromDomainTodo(t *todo.Todo) Todo {
	return Todo{
		ID:    t.ID,
		Title: t.Title,
	}
}
