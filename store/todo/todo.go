package todo

import (
	"github.com/google/uuid"

	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

type Todo struct {
	ID    uuid.UUID
	Title string
}

func NewTodo(id uuid.UUID, title string) *Todo {
	return &Todo{
		ID:    id,
		Title: title,
	}
}

func FromDomainTodo(t *todo.Todo) *Todo {
	return NewTodo(t.ID, t.Title)
}

func (t *Todo) ToDomainTodo() (*todo.Todo, error) {
	domainTodo, err := todo.NewTodo(t.ID, t.Title)

	return domainTodo, err
}
