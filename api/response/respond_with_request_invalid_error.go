package response

import (
	"log/slog"
	"net/http"
)

func RespondWithRequestInvalidError(detail string, w http.ResponseWriter, logger *slog.Logger) {
	problemDetails := NewProblemDetails().
		WithType("REQUEST_INVALID").
		WithStatus(http.StatusBadRequest).
		WithTitle("Invalid request").
		WithDetail(detail)

	RespondWithProblemDetails(problemDetails, w, logger)
}
