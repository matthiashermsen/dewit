package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func RespondWithProblemDetails(problemDetails ProblemDetails, w http.ResponseWriter, logger *slog.Logger) {
	bytes, err := json.Marshal(problemDetails)

	if err != nil {
		logger.Error("Could not marshal problem details", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.WriteHeader(problemDetails.Status)
	err = RespondWithBytes(w, bytes, logger)

	if err != nil {
		logger.Error("Could not write bytes", slog.Any("error", err))
	}
}
