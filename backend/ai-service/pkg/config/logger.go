package config

import (
	"log"
	"log/slog"
	"os"
)

func SetupLogger() *log.Logger {
	logHandler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.NewLogLogger(logHandler, slog.LevelError)
	slog.SetDefault(slog.New(logHandler))

	return logger
}
