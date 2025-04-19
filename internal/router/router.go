package router

import (
	"github.com/LuiggiPasacheL/go-url-shortener/internal/api"
	"github.com/LuiggiPasacheL/go-url-shortener/pkg/log"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h api.UrlHandler) *gin.Engine {

	r := gin.Default()

    r.Use(gin.Recovery())

    r.Use(log.LoggerMiddleware())

	r.GET("/health", h.Health)

	r.GET("/urls", h.GetAllUrls)

	r.POST("/urls", h.CreateUrl)

	r.GET("/r", h.RedirectUrl)

	r.GET("/urls/:id", h.GetUrl)

	return r
}

