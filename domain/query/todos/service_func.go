package todos

import "github.com/matthiashermsen/dewit/domain/entity/todo"

type ServiceFunc func() ([]*todo.Todo, error)
