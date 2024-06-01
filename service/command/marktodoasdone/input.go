package marktodoasdone

import "github.com/google/uuid"

type Input struct {
	TodoID uuid.UUID
}

func NewInput(todoID uuid.UUID) Input {
	return Input{
		TodoID: todoID,
	}
}
