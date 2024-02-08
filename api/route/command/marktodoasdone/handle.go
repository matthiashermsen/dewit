package marktodoasdone

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/request"
	"github.com/matthiashermsen/dewit/api/response"
	marktodoasdoneservice "github.com/matthiashermsen/dewit/domain/command/marktodoasdone"
	"github.com/matthiashermsen/dewit/domain/entity/todo"
)

func Handle(logger *slog.Logger, markTodoAsDone marktodoasdoneservice.ServiceFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody RequestBody
		err := request.ParseJSONBody(r, &requestBody)

		if err != nil {
			response.RespondWithJSONParseError(w, logger)

			return
		}

		input := marktodoasdoneservice.Input{TodoID: requestBody.TodoID}
		err = markTodoAsDone(input)

		if err != nil {
			_, isTodoNotFoundError := err.(*todo.TodoNotFoundError)

			if isTodoNotFoundError {
				problemDetails := response.
					NewProblemDetails().
					WithType("TODO_NOT_FOUND").
					WithStatus(http.StatusNotFound).
					WithTitle("Todo not found").
					WithDetail(fmt.Sprintf("Todo with ID '%s' wasn't found.", requestBody.TodoID))

				response.RespondWithProblemDetails(problemDetails, w, logger)

				return
			}

			logger.Error("Could not mark todo as done", slog.Any("error", err))
			response.RespondWithInternalServerError(w, logger)

			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
