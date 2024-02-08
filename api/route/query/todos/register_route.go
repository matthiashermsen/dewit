package todos

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/route/query"
	"github.com/matthiashermsen/dewit/domain/query/todos"
)

func RegisterRoute(mux *http.ServeMux, logger *slog.Logger, getTodos todos.ServiceFunc) {
	query.RegisterQueryRoute(mux, "todos", Handle(logger, getTodos))
}
