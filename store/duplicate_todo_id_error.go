package store

import (
	"fmt"

	"github.com/google/uuid"
)

type DuplicateTodoIDError struct {
	todoID uuid.UUID
}

func NewDuplicateTodoIDError(todoID uuid.UUID) *DuplicateTodoIDError {
	return &DuplicateTodoIDError{
		todoID: todoID,
	}
}

func (e *DuplicateTodoIDError) Error() string {
	return fmt.Sprintf("Duplicate todo ID '%s'.", e.todoID)
}
