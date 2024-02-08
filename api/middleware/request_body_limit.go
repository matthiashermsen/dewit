package middleware

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/response"
)

func RequestBodyLimit(logger *slog.Logger, maximumBytes int64, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ContentLength > maximumBytes {
			problemDetails := response.
				NewProblemDetails().
				WithType("CONTENT_TOO_LARGE").
				WithStatus(http.StatusRequestEntityTooLarge).
				WithTitle("Body size exceeding limit").
				WithDetail(fmt.Sprintf("Body size is exceeding the limit of '%d' bytes", maximumBytes))

			response.RespondWithProblemDetails(problemDetails, w, logger)

			return
		}

		next.ServeHTTP(w, r)
	})
}
