package store

import (
	"fmt"

	"github.com/google/uuid"
)

type TodoNotFoundError struct {
	todoID uuid.UUID
}

func NewTodoNotFoundError(todoID uuid.UUID) *TodoNotFoundError {
	return &TodoNotFoundError{
		todoID: todoID,
	}
}

func (e *TodoNotFoundError) Error() string {
	return fmt.Sprintf("Todo with ID '%s' not found.", e.todoID)
}
