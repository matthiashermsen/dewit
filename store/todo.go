package store

import (
	"github.com/google/uuid"

	"github.com/matthiashermsen/dewit/domain"
)

type Todo struct {
	id             uuid.UUID
	title          string
	isMarkedAsDone bool
}

func NewTodo(id uuid.UUID, title string, isMarkedAsDone bool) *Todo {
	return &Todo{
		id:             id,
		title:          title,
		isMarkedAsDone: isMarkedAsDone,
	}
}

func (todo *Todo) GetID() uuid.UUID {
	return todo.id
}

func (todo *Todo) GetTitle() string {
	return todo.title
}

func (todo *Todo) IsMarkedAsDone() bool {
	return todo.isMarkedAsDone
}

func (todo *Todo) MarkAsDone() {
	todo.isMarkedAsDone = true
}

func FromDomainTodo(t *domain.Todo) *Todo {
	return NewTodo(t.GetID(), t.GetTitle(), t.IsMarkedAsDone())
}

func (t *Todo) ToDomainTodo() (*domain.Todo, error) {
	domainTodo, err := domain.NewTodo(t.id, t.title)

	if err != nil {
		return nil, err
	}

	if t.isMarkedAsDone {
		err = domainTodo.MarkAsDone()
	}

	return domainTodo, err
}
