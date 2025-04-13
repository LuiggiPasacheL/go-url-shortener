package main

import (
	"github.com/LuiggiPasacheL/go-url-shortener/internal/router"
	"github.com/LuiggiPasacheL/go-url-shortener/pkg/log"
)


func main() {

    logger := log.SetupLogger()

    logger.Info("Starting URL shortener service")

	r := router.SetupRouter()

    if err := r.Run(":8080"); err != nil {
        logger.Error("Failed to start server", "error", err)
    } else {
        logger.Info("Server started successfully", "port", ":8080")
    }
}


