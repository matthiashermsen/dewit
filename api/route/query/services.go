package query

import (
	"github.com/matthiashermsen/dewit/domain/query/todos"
)

type Services struct {
	GetTodos todos.ServiceFunc
}
