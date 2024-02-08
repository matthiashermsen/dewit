package todos

import "github.com/matthiashermsen/dewit/store/todo"

type TodoStore interface {
	GetTodos() ([]*todo.Todo, error)
}
