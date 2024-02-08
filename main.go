package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/matthiashermsen/dewit/api"
	"github.com/matthiashermsen/dewit/api/route/command"
	"github.com/matthiashermsen/dewit/api/route/query"
	"github.com/matthiashermsen/dewit/appconfiguration"
	"github.com/matthiashermsen/dewit/appversion"
	"github.com/matthiashermsen/dewit/domain/command/marktodoasdone"
	"github.com/matthiashermsen/dewit/domain/command/notetodo"
	"github.com/matthiashermsen/dewit/domain/query/todos"
	"github.com/matthiashermsen/dewit/logging"
	"github.com/matthiashermsen/dewit/store/todo"
)

func main() {
	configuration, err := appconfiguration.New()

	if err != nil {
		panic(err)
	}

	logger := logging.GetLogger(configuration.Logging.Level)
	todoStore := todo.NewInMemoryStore(logger)

	getTodosService := todos.Handle(logger, todoStore)
	queryServices := query.Services{
		GetTodos: getTodosService,
	}

	noteTodoService := notetodo.Handle(logger, todoStore)
	markTodoAsDoneService := marktodoasdone.Handle(logger, todoStore)
	commandServices := command.Services{
		NoteTodo:       noteTodoService,
		MarkTodoAsDone: markTodoAsDoneService,
	}

	api := api.GetAPI(logger, appversion.AppVersion, queryServices, commandServices)

	logger.Info(fmt.Sprintf("Starting server on port %d", configuration.API.Port))

	err = http.ListenAndServe(":"+strconv.Itoa(int(configuration.API.Port)), api)

	if err != nil {
		logger.Error("Could not serve API", slog.Any("error", err))

		os.Exit(1)
	}
}
