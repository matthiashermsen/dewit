package api

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/route/command"
	"github.com/matthiashermsen/dewit/api/route/command/marktodoasdone"
	"github.com/matthiashermsen/dewit/api/route/command/notetodo"
	"github.com/matthiashermsen/dewit/api/route/getappversion"
	"github.com/matthiashermsen/dewit/api/route/notfound"
	"github.com/matthiashermsen/dewit/api/route/ping"
	"github.com/matthiashermsen/dewit/api/route/query"
	"github.com/matthiashermsen/dewit/api/route/query/todos"
)

func GetAPI(logger *slog.Logger, appVersion string, queryServices query.Services, commandServices command.Services) *http.ServeMux {
	mux := http.NewServeMux()

	// queries
	todos.RegisterRoute(mux, logger, queryServices.GetTodos)

	// commands
	notetodo.RegisterRoute(mux, logger, commandServices.NoteTodo)
	marktodoasdone.RegisterRoute(mux, logger, commandServices.MarkTodoAsDone)

	// technical routes
	ping.RegisterRoute(mux)
	getappversion.RegisterRoute(mux, logger, appVersion)

	notfound.RegisterRoute(mux, logger)

	return mux
}
