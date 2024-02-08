package marktodoasdone

import (
	"log/slog"
	"math"
	"net/http"

	"github.com/matthiashermsen/dewit/api/middleware"
	"github.com/matthiashermsen/dewit/api/route/command"
	"github.com/matthiashermsen/dewit/domain/command/marktodoasdone"
)

func RegisterRoute(mux *http.ServeMux, logger *slog.Logger, markTodoAsDone marktodoasdone.ServiceFunc) {
	maximumBodySizeInBytes := int64(math.MaxInt64)

	command.RegisterCommandRoute(mux, "mark-todo-as-done", middleware.RequestBodyLimit(logger, maximumBodySizeInBytes, Handle(logger, markTodoAsDone)))
}
