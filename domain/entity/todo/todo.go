package todo

import "github.com/google/uuid"

type Todo struct {
	ID    uuid.UUID
	Title string
}

func NewTodo(id uuid.UUID, title string) (*Todo, error) {
	if title == "" {
		return nil, &TodoTitleEmptyError{}
	}

	todo := Todo{
		ID:    id,
		Title: title,
	}

	return &todo, nil
}
