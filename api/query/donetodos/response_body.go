package donetodos

import (
	"github.com/google/uuid"

	"github.com/matthiashermsen/dewit/domain"
)

type ResponseBody []todoListItem

func FromDomainTodos(domainTodos []*domain.Todo) ResponseBody {
	todoListItems := make([]todoListItem, 0)

	for _, domainTodo := range domainTodos {
		todoListItem := newTodoListItem(domainTodo.GetID(), domainTodo.GetTitle())
		todoListItems = append(todoListItems, todoListItem)
	}

	return todoListItems
}

type todoListItem struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

func newTodoListItem(id uuid.UUID, title string) todoListItem {
	return todoListItem{
		ID:    id,
		Title: title,
	}
}
