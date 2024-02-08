package notetodo

import "github.com/matthiashermsen/dewit/domain/entity/todo"

type ServiceFunc func(i Input) (*todo.Todo, error)
