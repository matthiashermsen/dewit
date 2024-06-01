package donetodos

import "github.com/matthiashermsen/dewit/store"

type Store interface {
	GetDoneTodos() ([]*store.Todo, error)
}
