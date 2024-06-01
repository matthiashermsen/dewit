package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type TodoAlreadyMarkedAsDoneError struct {
	todoID uuid.UUID
}

func NewTodoAlreadyMarkedAsDoneError(todoID uuid.UUID) *TodoAlreadyMarkedAsDoneError {
	return &TodoAlreadyMarkedAsDoneError{
		todoID: todoID,
	}
}

func (e *TodoAlreadyMarkedAsDoneError) Error() string {
	return fmt.Sprintf("Todo with ID '%s' is already marked as done.", e.todoID)
}
