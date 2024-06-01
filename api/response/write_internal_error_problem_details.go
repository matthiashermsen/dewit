package response

import (
	"log/slog"
	"net/http"
)

func WriteInternalErrorProblemDetails(w http.ResponseWriter, logger *slog.Logger) {
	problemDetails := NewProblemDetails("INTERNAL_ERROR", http.StatusInternalServerError, "Internal error", "The server encountered an unexpected condition that prevented it from fulfilling the request.")

	WriteProblemDetails(w, problemDetails, logger)
}
