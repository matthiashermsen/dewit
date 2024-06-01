package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/matthiashermsen/dewit/api"
	"github.com/matthiashermsen/dewit/app"
	"github.com/matthiashermsen/dewit/cfg"
	"github.com/matthiashermsen/dewit/log"
	"github.com/matthiashermsen/dewit/store"
)

func main() {
	logLevel, err := cfg.GetLogLevel()

	if err != nil {
		panic(err)
	}

	logger := log.New(logLevel)
	inMemoryStore := store.NewInMemoryStore(logger)

	mux := api.New(app.Version, inMemoryStore, logger)

	port := cfg.GetPort()
	address := fmt.Sprintf(":%s", port)

	logger.Info(fmt.Sprintf("Running on version %s", app.Version))
	logger.Info(fmt.Sprintf("Server listening on port %s", port))

	err = http.ListenAndServe(address, mux)

	if err != nil {
		log.Error(logger, "Could not serve API", err)

		os.Exit(1)
	}
}
