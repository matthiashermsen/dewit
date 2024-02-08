package todo

import (
	"log/slog"
	"sync"

	"github.com/google/uuid"
	"github.com/matthiashermsen/dewit/store"
)

type InMemoryStore struct {
	logger *slog.Logger
	todos  []*Todo
	mutex  sync.RWMutex
}

func NewInMemoryStore(logger *slog.Logger) *InMemoryStore {
	return &InMemoryStore{
		logger: logger,
	}
}

func (s *InMemoryStore) GetTodos() ([]*Todo, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	todos := make([]Todo, 0)

	for _, todo := range s.todos {
		todos = append(todos, *todo)
	}

	return s.todos, nil
}

func (s *InMemoryStore) SaveTodo(todo *Todo) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, currentTodo := range s.todos {
		if currentTodo.ID == todo.ID {
			return store.NewDuplicateEntityIdError(todo.ID.String())
		}
	}

	s.todos = append(s.todos, todo)

	return nil
}

func (s *InMemoryStore) DeleteTodoById(todoID uuid.UUID) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for index, todo := range s.todos {
		if todo.ID == todoID {
			s.todos = append(s.todos[:index], s.todos[index+1:]...)

			return nil
		}
	}

	return store.NewEntityNotFoundError(todoID.String())
}
