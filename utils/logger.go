package utils

import (
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
	programLevel *slog.LevelVar
}

var logger *Logger

func InitLogger(level slog.Level) {
	programLevel := new(slog.LevelVar)
	l := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	programLevel.Set(level)
	logger = &Logger{
		Logger:       slog.New(l),
		programLevel: programLevel,
	}
	slog.SetDefault(logger.Logger)
}

func SetLogLevel(level slog.Level) {
	if logger == nil {
		InitLogger(level)
	}
	logger.programLevel.Set(level)
}
