package notfound

import (
	"log/slog"
	"net/http"
)

func RegisterRoute(mux *http.ServeMux, logger *slog.Logger) {
	mux.HandleFunc("/", Handle(logger))
}
