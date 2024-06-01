package store

import (
	"log/slog"
	"sync"

	"github.com/google/uuid"
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

func (s *InMemoryStore) GetRemainingTodos() ([]*Todo, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	todos := make([]*Todo, 0)

	for _, todo := range s.todos {
		if todo.IsMarkedAsDone() {
			continue
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *InMemoryStore) GetDoneTodos() ([]*Todo, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	todos := make([]*Todo, 0)

	for _, todo := range s.todos {
		if !todo.IsMarkedAsDone() {
			continue
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (s *InMemoryStore) GetTodoByID(todoID uuid.UUID) (*Todo, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, currentTodo := range s.todos {
		if currentTodo.GetID() == todoID {
			return currentTodo, nil
		}
	}

	return nil, NewTodoNotFoundError(todoID)
}

func (s *InMemoryStore) CreateTodo(todo *Todo) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, currentTodo := range s.todos {
		todoID := todo.GetID()

		if currentTodo.GetID() == todoID {
			return NewDuplicateTodoIDError(todoID)
		}
	}

	s.todos = append(s.todos, todo)

	return nil
}

func (s *InMemoryStore) MarkTodoAsDone(todoID uuid.UUID) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, currentTodo := range s.todos {
		if currentTodo.GetID() == todoID {
			currentTodo.MarkAsDone()

			return nil
		}
	}

	return NewTodoNotFoundError(todoID)
}
