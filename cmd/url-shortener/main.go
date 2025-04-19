package main

import (
	"os"

	"github.com/LuiggiPasacheL/go-url-shortener/internal/api"
	"github.com/LuiggiPasacheL/go-url-shortener/internal/repository"
	"github.com/LuiggiPasacheL/go-url-shortener/internal/router"
	"github.com/LuiggiPasacheL/go-url-shortener/internal/services"
	"github.com/LuiggiPasacheL/go-url-shortener/pkg/log"
)


func main() {

    logger := log.SetupLogger()

    logger.Info("Starting URL shortener service")

	repository, err := repository.NewUrlRepositorySqlite("db.sqlite")
	if err != nil {
		logger.Error("Failed to start repository", "error", err)
		os.Exit(1)
	}
	defer repository.Close()

	service := services.NewUrlServiceImpl(repository)

	handler := api.NewUrlHandler(logger, service)

	r := router.SetupRouter(handler)

    if err := r.Run(":8080"); err != nil {
        logger.Error("Failed to start server", "error", err)
    } else {
        logger.Info("Server started successfully", "port", ":8080")
    }
}


