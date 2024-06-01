package remainingtodos

import "github.com/matthiashermsen/dewit/store"

type Store interface {
	GetRemainingTodos() ([]*store.Todo, error)
}
