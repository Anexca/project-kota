package config

import (
	"context"
	"log"
	"os"
	"time"

	"golang.org/x/exp/slog"
)

// color codes for different log levels
const (
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorReset  = "\033[0m"
)

// colorfulHandler is a custom handler to add color and formatting
type colorfulHandler struct {
	slog.Handler
}

// Handle method to customize the log output
func (h *colorfulHandler) Handle(ctx context.Context, r slog.Record) error {
	var levelColor string

	switch r.Level {
	case slog.LevelError:
		levelColor = colorRed
	case slog.LevelWarn:
		levelColor = colorYellow
	case slog.LevelInfo:
		levelColor = colorGreen
	case slog.LevelDebug:
		levelColor = colorBlue
	default:
		levelColor = colorReset
	}

	// Construct the log message
	message := levelColor + "[" + r.Level.String() + "]" + colorReset + " " + r.Time.Format(time.RFC3339) + " " + r.Message

	r.Attrs(func(attr slog.Attr) bool {
		message += " " + attr.Key + "=" + attr.Value.String()
		return true
	})

	_, err := os.Stdout.WriteString(message + "\n")
	return err
}

// SetupLogger initializes and returns a colorful, informative logger
func SetupLogger() *log.Logger {
	logHandler := &colorfulHandler{
		Handler: slog.NewJSONHandler(os.Stdout, nil),
	}

	// Setting up the default logger with custom handler and error level
	logger := slog.NewLogLogger(logHandler, slog.LevelInfo)
	slog.SetDefault(slog.New(logHandler))

	return logger
}
