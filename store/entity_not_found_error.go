package store

import "fmt"

type EntityNotFoundError struct {
	entityID string
}

func NewEntityNotFoundError(entityID string) *EntityNotFoundError {
	return &EntityNotFoundError{
		entityID: entityID,
	}
}

func (e *EntityNotFoundError) Error() string {
	return fmt.Sprintf("Entity with id '%s' not found.", e.entityID)
}
