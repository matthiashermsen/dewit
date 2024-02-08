package notetodo

import (
	"log/slog"
	"math"
	"net/http"

	"github.com/matthiashermsen/dewit/api/middleware"
	"github.com/matthiashermsen/dewit/api/route/command"
	"github.com/matthiashermsen/dewit/domain/command/notetodo"
)

func RegisterRoute(mux *http.ServeMux, logger *slog.Logger, noteTodo notetodo.ServiceFunc) {
	maximumBodySizeInBytes := int64(math.MaxInt64)

	command.RegisterCommandRoute(mux, "note-todo", middleware.RequestBodyLimit(logger, maximumBodySizeInBytes, Handle(logger, noteTodo)))
}
