package query

import "github.com/matthiashermsen/dewit/store"

type Store interface {
	GetRemainingTodos() ([]*store.Todo, error)
	GetDoneTodos() ([]*store.Todo, error)
}
