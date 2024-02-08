package marktodoasdone

import "github.com/google/uuid"

type RequestBody struct {
	TodoID uuid.UUID `json:"todoID"`
}
