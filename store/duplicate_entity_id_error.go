package store

import "fmt"

type DuplicateEntityIdError struct {
	entityID string
}

func NewDuplicateEntityIdError(entityID string) *DuplicateEntityIdError {
	return &DuplicateEntityIdError{
		entityID: entityID,
	}
}

func (e *DuplicateEntityIdError) Error() string {
	return fmt.Sprintf("Duplicate ID '%s' for entity.", e.entityID)
}
