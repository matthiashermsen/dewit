package command

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/command/marktodoasdone"
	"github.com/matthiashermsen/dewit/api/command/notetodo"
	"github.com/matthiashermsen/dewit/api/route"
)

func New(store Store, logger *slog.Logger) *http.ServeMux {
	api := http.NewServeMux()

	api.HandleFunc(GetRoutePattern("note-todo"), route.ChainHandlers(notetodo.Handle(store, logger)))
	api.HandleFunc(GetRoutePattern("mark-todo-as-done"), route.ChainHandlers(marktodoasdone.Handle(store, logger)))

	return api
}
