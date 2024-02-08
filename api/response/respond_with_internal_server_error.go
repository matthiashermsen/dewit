package response

import (
	"log/slog"
	"net/http"
)

func RespondWithInternalServerError(w http.ResponseWriter, logger *slog.Logger) {
	problemDetails := GetInternalServerErrorProblemDetails()

	RespondWithProblemDetails(problemDetails, w, logger)
}
