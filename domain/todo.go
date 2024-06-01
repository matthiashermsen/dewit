package domain

import "github.com/google/uuid"

type Todo struct {
	id             uuid.UUID
	title          string
	isMarkedAsDone bool
}

func NewTodo(id uuid.UUID, title string) (*Todo, error) {
	if title == "" {
		return nil, ErrTodoTitleEmpty
	}

	todo := Todo{
		id:             id,
		title:          title,
		isMarkedAsDone: false,
	}

	return &todo, nil
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

func (todo *Todo) MarkAsDone() error {
	if todo.isMarkedAsDone {
		err := NewTodoAlreadyMarkedAsDoneError(todo.id)

		return err
	}

	todo.isMarkedAsDone = true

	return nil
}
