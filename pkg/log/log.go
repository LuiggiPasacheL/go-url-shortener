package log

import (
	"context"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
)

var LoggerKey string = "logger"
var RequestIDKey string = "request_id"
var MethodKey string = "method"
var PathKey string = "path"
var IPKey string = "ip"

func SetupLogger() *slog.Logger {
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return logger
}

func GetLogger(c context.Context) *slog.Logger {
	logger, ok := c.Value(LoggerKey).(*slog.Logger)
	if !ok {
		logger = SetupLogger()
	}
    
    requestID, ok := c.Value(RequestIDKey).(string)
    if ok {
        logger = logger.With(RequestIDKey, requestID)
    }
    method, ok := c.Value(MethodKey).(string)
    if ok {
        logger = logger.With(MethodKey, method)
    }
    path, ok := c.Value(PathKey).(string)
    if ok {
        logger = logger.With(PathKey, path)
    }
    ip, ok := c.Value(IPKey).(string)
    if ok {
        logger = logger.With(IPKey, ip)
    }

    return logger
}

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        logger := SetupLogger()
        c.Set(LoggerKey, logger)
        c.Set(RequestIDKey, c.Request.Header.Get("X-Request-ID"))
        c.Set(MethodKey, c.Request.Method)
        c.Set(PathKey, c.Request.URL.Path)
        c.Set(IPKey, c.ClientIP())
        c.Next()
    }
}
