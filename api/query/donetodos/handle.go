package donetodos

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/response"
	logging "github.com/matthiashermsen/dewit/log"
	donetodosservice "github.com/matthiashermsen/dewit/service/query/donetodos"
)

func Handle(dataStore donetodosservice.Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := donetodosservice.Handle(dataStore)

		if err != nil {
			logging.Error(logger, err.Error(), err)

			response.WriteInternalErrorProblemDetails(w, logger)

			return
		}

		w.WriteHeader(http.StatusOK)
		responseBody := FromDomainTodos(todos)
		response.WriteJSON(w, responseBody, logger)
	}
}
