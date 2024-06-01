package response

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/log"
)

func WriteProblemDetails(w http.ResponseWriter, problemDetails ProblemDetails, logger *slog.Logger) {
	bytes, err := json.Marshal(problemDetails)

	if err != nil {
		log.Error(logger, "Could not marshal problem details", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Header().Add("Content-Type", "application/problem+json")
	w.WriteHeader(problemDetails.Status)
	_, err = w.Write(bytes)

	if err != nil {
		log.Error(logger, "Could not write bytes", err)
	}
}
