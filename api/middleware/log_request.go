package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedResponseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LogRequest(logger *slog.Logger, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrappedResponseWriter := &wrappedResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrappedResponseWriter, r)

		logger.Debug(
			"Request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("statusCode", wrappedResponseWriter.statusCode),
			slog.Duration("duration", time.Since(start)),
		)
	})
}
