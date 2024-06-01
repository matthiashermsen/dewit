package marktodoasdone

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/matthiashermsen/dewit/api/request"
	"github.com/matthiashermsen/dewit/api/response"
	"github.com/matthiashermsen/dewit/domain"
	logging "github.com/matthiashermsen/dewit/log"
	marktodoasdoneservice "github.com/matthiashermsen/dewit/service/command/marktodoasdone"
	"github.com/matthiashermsen/dewit/store"
)

func Handle(dataStore marktodoasdoneservice.Store, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody RequestBody

		err := request.ParseJSON(r, &requestBody)

		if err != nil {
			response.WriteRequestBodyInvalidProblemDetails(w, err, logger)

			return
		}

		input := marktodoasdoneservice.NewInput(requestBody.TodoID)
		err = marktodoasdoneservice.Handle(input, dataStore)

		var todoNotFoundError *store.TodoNotFoundError

		if errors.As(err, &todoNotFoundError) {
			problemDetails := response.NewProblemDetails("TODO_NOT_FOUND", http.StatusNotFound, "Todo not found", fmt.Sprintf("Could not find todo with ID '%s'.", requestBody.TodoID))
			response.WriteProblemDetails(w, problemDetails, logger)

			return
		}

		var todoAlreadyMarkedAsDoneError *domain.TodoAlreadyMarkedAsDoneError

		if errors.As(err, &todoAlreadyMarkedAsDoneError) {
			problemDetails := response.NewProblemDetails("TODO_ALREADY_MARKED_AS_DONE", http.StatusConflict, "Todo already marked as done", fmt.Sprintf("Todo with ID '%s' is already marked as done.", requestBody.TodoID))
			response.WriteProblemDetails(w, problemDetails, logger)

			return
		}

		if err != nil {
			logging.Error(logger, err.Error(), err)

			response.WriteInternalErrorProblemDetails(w, logger)

			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
