package notetodo

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/request"
	"github.com/matthiashermsen/dewit/api/response"
	"github.com/matthiashermsen/dewit/api/responsetype"
	notetodoservice "github.com/matthiashermsen/dewit/domain/command/notetodo"
	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

func Handle(logger *slog.Logger, noteTodo notetodoservice.ServiceFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody RequestBody
		err := request.ParseJSONBody(r, &requestBody)

		if err != nil {
			response.RespondWithJSONParseError(w, logger)

			return
		}

		input := notetodoservice.Input{TodoTitle: requestBody.TodoTitle}
		notedTodo, err := noteTodo(input)

		if err != nil {
			_, isTodoTitleEmptyError := err.(*todo.TodoTitleEmptyError)

			if isTodoTitleEmptyError {
				problemDetails := response.
					NewProblemDetails().
					WithType("TODO_TITLE_EMPTY").
					WithStatus(http.StatusBadRequest).
					WithTitle("Todo title is empty").
					WithDetail("Todo title must not be empty.")

				response.RespondWithProblemDetails(problemDetails, w, logger)

				return
			}

			logger.Error("Could not note todo", slog.Any("error", err))
			response.RespondWithInternalServerError(w, logger)

			return
		}

		responseTodo := responsetype.FromDomainTodo(notedTodo)
		bytes, err := json.Marshal(responseTodo)

		if err != nil {
			logger.Error("Could not marshal todo", slog.Any("error", err))
			response.RespondWithInternalServerError(w, logger)

			return
		}

		response.RespondWithBytes(w, bytes, logger)
	}
}
