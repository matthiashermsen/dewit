package notetodo

import (
	"github.com/google/uuid"

	"github.com/matthiashermsen/dewit/domain"
)

type ResponseBody struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

func FromDomainTodo(todo *domain.Todo) ResponseBody {
	return ResponseBody{
		ID:    todo.GetID(),
		Title: todo.GetTitle(),
	}
}
