package todos

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/response"
	"github.com/matthiashermsen/dewit/api/responsetype"
	todosservice "github.com/matthiashermsen/dewit/domain/query/todos"
)

func Handle(logger *slog.Logger, getTodos todosservice.ServiceFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := getTodos()

		if err != nil {
			logger.Error("Could not retrieve todos", slog.Any("error", err))
			response.RespondWithInternalServerError(w, logger)

			return
		}

		responseTodos := make([]responsetype.Todo, 0)

		for _, todo := range todos {
			responseTodo := responsetype.FromDomainTodo(todo)
			responseTodos = append(responseTodos, responseTodo)
		}

		response.RespondWithJSON(w, responseTodos, logger)
	}
}
