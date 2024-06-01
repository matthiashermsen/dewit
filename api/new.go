package api

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/appversion"
	"github.com/matthiashermsen/dewit/api/command"
	"github.com/matthiashermsen/dewit/api/middleware"
	"github.com/matthiashermsen/dewit/api/notfound"
	"github.com/matthiashermsen/dewit/api/ping"
	"github.com/matthiashermsen/dewit/api/query"
)

func New(appVersion string, store Store, logger *slog.Logger) http.Handler {
	api := http.NewServeMux()

	api.HandleFunc("GET /ping", ping.Handle())
	api.HandleFunc("GET /app-version", appversion.Handle(appVersion, logger))

	queryRouter := query.New(store, logger)
	api.Handle("/query/", http.StripPrefix("/query", queryRouter))

	commandRouter := command.New(store, logger)
	api.Handle("/command/", http.StripPrefix("/command", commandRouter))

	api.HandleFunc("/", notfound.Handle(logger))

	return middleware.LogRequest(logger, api.ServeHTTP)
}
