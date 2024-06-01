package notetodo

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/request"
	"github.com/matthiashermsen/dewit/api/response"
	"github.com/matthiashermsen/dewit/domain"
	logging "github.com/matthiashermsen/dewit/log"
	notetodoservice "github.com/matthiashermsen/dewit/service/command/notetodo"
)

func Handle(store notetodoservice.Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody RequestBody

		err := request.ParseJSON(r, &requestBody)

		if err != nil {
			response.WriteRequestBodyInvalidProblemDetails(w, err, logger)

			return
		}

		input := notetodoservice.NewInput(requestBody.TodoTitle)
		todo, err := notetodoservice.Handle(input, store)

		if errors.Is(err, domain.ErrTodoTitleEmpty) {
			problemDetails := response.NewProblemDetails("TODO_TITLE_EMPTY", http.StatusBadRequest, "Todo title is empty", "Todo title must not be empty.")
			response.WriteProblemDetails(w, problemDetails, logger)

			return
		}

		if err != nil {
			logging.Error(logger, err.Error(), err)

			response.WriteInternalErrorProblemDetails(w, logger)

			return
		}

		w.WriteHeader(http.StatusCreated)

		responseBody := FromDomainTodo(todo)
		response.WriteJSON(w, responseBody, logger)
	}
}
