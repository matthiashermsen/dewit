package response

import (
	"log/slog"
	"net/http"
)

func WriteRequestBodyInvalidProblemDetails(w http.ResponseWriter, err error, logger *slog.Logger) {
	problemDetails := NewProblemDetails("REQUEST_BODY_INVALID", http.StatusBadRequest, "Request body invalid", err.Error())

	WriteProblemDetails(w, problemDetails, logger)
}
