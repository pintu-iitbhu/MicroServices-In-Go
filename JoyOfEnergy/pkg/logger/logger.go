package logger

import (
	"golang.org/x/exp/slog"
	"os"
)

var logger *slog.Logger

func NewLogger() *slog.Logger {
	if logger != nil {
		return logger
	}
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return logger
}
