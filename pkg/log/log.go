package log

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func SetupLogger() *slog.Logger {

    if logger != nil {
        return logger
    }

    logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
    return logger
}
