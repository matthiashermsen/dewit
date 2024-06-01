package notfound

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/response"
)

func Handle(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		problemDetails := response.ProblemDetails{
			Type:   "RESOURCE_MISSING",
			Status: http.StatusNotFound,
			Title:  "Not found",
			Detail: fmt.Sprintf("Could not find %s '%s'.", r.Method, r.URL.Path),
		}

		response.WriteProblemDetails(w, problemDetails, logger)
	}
}
