package main

import (
	"log/slog"
	"strings"
	"fmt"
	"os"
)

func setupLogger(levelFromConfig string) {
	level := getLevel(levelFromConfig)
	handler := getHandler(level)
	setHandlerToDefaultLogger(handler)
}

func getLevel(levelFromConfig string) slog.Level {
	levelFromConfig = strings.ToLower(levelFromConfig)
	var level slog.Level
	if err := level.UnmarshalText([]byte(levelFromConfig)); err != nil {
		level = slog.LevelInfo
		slog.Warn(fmt.Sprintf("Invalid value (%s) of LOG_LEVEL in config. setup level to info", levelFromConfig))
	}

	return level
}

func getHandler(level slog.Level) slog.Handler {
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level: level,
	})
}

func setHandlerToDefaultLogger(handler slog.Handler) {
	newLogger := slog.New(handler)
	slog.SetDefault(newLogger)
}
