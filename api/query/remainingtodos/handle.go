package remainingtodos

import (
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/response"
	logging "github.com/matthiashermsen/dewit/log"
	remainingtodosservice "github.com/matthiashermsen/dewit/service/query/remainingtodos"
)

func Handle(dataStore remainingtodosservice.Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		todos, err := remainingtodosservice.Handle(dataStore)

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
