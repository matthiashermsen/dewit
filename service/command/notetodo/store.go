package notetodo

import (
	"github.com/matthiashermsen/dewit/store"
)

type Store interface {
	CreateTodo(todo *store.Todo) error
}
