package router

import (
	"net/http"

	"github.com/LuiggiPasacheL/go-url-shortener/pkg/log"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()

    r.Use(gin.Recovery())

    r.Use(log.LoggerMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

