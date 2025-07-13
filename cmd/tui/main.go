package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/nifle3/tui_music/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

var version string

func main() {
	config := mustNewConfig()
	setupLogger(config.LogLevel)

	slog.Debug(fmt.Sprintf("Init program with config %#v", config))
	slog.Info("Logger setup")

	_ = flag.String("yandex_oauth", "", "Yandex OAUTH key to access music api via account")
	flag.Parse()

	slog.Info("Application doesn't run, starting")

	p := tea.NewProgram(ui.NewApp(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		slog.Error("There here is an error", slog.String("Error", err.Error()))
	}
}
