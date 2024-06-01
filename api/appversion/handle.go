package appversion

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/response"
)

func Handle(appVersion string, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if appVersion == "" {
			problemDetails := response.ProblemDetails{
				Type:   "APP_VERSION_UNAVAILABLE",
				Status: http.StatusNotFound,
				Title:  "App version is unavailable",
				Detail: "Current app version is not available.",
			}

			response.WriteProblemDetails(w, problemDetails, logger)

			return
		}

		response.WriteJSON(w, appVersion, logger)
	}
}
