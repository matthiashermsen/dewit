package getappversion

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/response"
)

func Handle(logger *slog.Logger, appVersion string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if appVersion == "" {
			problemDetails := response.
				NewProblemDetails().
				WithType("APP_VERSION_UNAVAILABLE").
				WithStatus(http.StatusNotFound).
				WithTitle("App version is unavailable").
				WithDetail("Current app version is not available.")

			response.RespondWithProblemDetails(problemDetails, w, logger)

			return
		}

		bytes, err := json.Marshal(appVersion)

		if err != nil {
			logger.Error("Could not marshal app version", slog.Any("error", err))
			response.RespondWithInternalServerError(w, logger)

			return
		}

		response.RespondWithBytes(w, bytes, logger)
	}
}
