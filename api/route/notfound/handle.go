package notfound

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/response"
)

func Handle(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		problemDetails := response.
			NewProblemDetails().
			WithType("RESOURCE_MISSING").
			WithStatus(http.StatusNotFound).
			WithTitle("Not found").
			WithDetail(fmt.Sprintf("Could not find '%s'", r.URL.Path))

		response.RespondWithProblemDetails(problemDetails, w, logger)
	}
}
