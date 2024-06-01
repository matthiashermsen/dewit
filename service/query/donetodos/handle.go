package donetodos

import "github.com/matthiashermsen/dewit/domain"

func Handle(dataStore Store) ([]*domain.Todo, error) {
	storeTodos, err := dataStore.GetDoneTodos()

	if err != nil {
		return nil, err
	}

	domainTodos := make([]*domain.Todo, 0)

	for _, storeTodo := range storeTodos {
		domainTodo, err := storeTodo.ToDomainTodo()

		if err != nil {
			return nil, err
		}

		domainTodos = append(domainTodos, domainTodo)
	}

	return domainTodos, nil
}
