package query

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/query/donetodos"
	"github.com/matthiashermsen/dewit/api/query/remainingtodos"
	"github.com/matthiashermsen/dewit/api/route"
)

func New(store Store, logger *slog.Logger) *http.ServeMux {
	api := http.NewServeMux()

	api.HandleFunc(GetRoutePattern("remaining-todos"), route.ChainHandlers(remainingtodos.Handle(store, logger)))
	api.HandleFunc(GetRoutePattern("done-todos"), route.ChainHandlers(donetodos.Handle(store, logger)))

	return api
}
