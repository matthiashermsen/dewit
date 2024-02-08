package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, payload any, logger *slog.Logger) {
	bytes, err := json.Marshal(payload)

	if err != nil {
		logger.Error("Could not encode payload", slog.Any("error", err))
		RespondWithInternalServerError(w, logger)

		return
	}

	err = RespondWithBytes(w, bytes, logger)

	if err != nil {
		logger.Error("Could not write bytes", slog.Any("error", err))
	}
}
