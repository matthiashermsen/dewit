package response

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/log"
)

func WriteJSON(w http.ResponseWriter, data any, logger *slog.Logger) {
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Error(logger, "Could not write data", err)
	}
}
