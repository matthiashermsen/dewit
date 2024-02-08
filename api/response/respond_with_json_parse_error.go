package response

import (
	"log/slog"
	"net/http"
)

func RespondWithJSONParseError(w http.ResponseWriter, logger *slog.Logger) {
	problemDetails := GetJSONParseErrorProblemDetails()

	RespondWithProblemDetails(problemDetails, w, logger)
}
