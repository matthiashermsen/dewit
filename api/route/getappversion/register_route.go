package getappversion

import (
	"log/slog"
	"net/http"
)

func RegisterRoute(mux *http.ServeMux, logger *slog.Logger, appVersion string) {
	mux.HandleFunc("/app-version", Handle(logger, appVersion))
}
