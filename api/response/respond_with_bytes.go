package response

import (
	"fmt"
	"log/slog"
	"net/http"
)

func RespondWithBytes(w http.ResponseWriter, bytes []byte, logger *slog.Logger) error {
	_, err := w.Write(bytes)

	if err != nil {
		return fmt.Errorf("Failed to write bytes: %w", err)
	}

	return nil
}
