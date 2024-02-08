package notetodo

import "github.com/matthiashermsen/dewit/store/todo"

type TodoStore interface {
	SaveTodo(todo *todo.Todo) error
}
