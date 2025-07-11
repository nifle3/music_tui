package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"sync"

	"github.com/nifle3/tui_music/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nifle3/tui_music/internal/ipc"
)

var version string

func main() {
	config := mustNewConfig()
	setupLogger(config.LogLevel)

	slog.Debug(fmt.Sprintf("Init program with config %#v", config))
	slog.Info("Logger setup")

	yandexOAUTHToken := flag.String("yandex_oauth", "", "Yandex OAUTH key to access music api via account")
	flag.Parse()

	if isSendFlagsToClientSuccess(*yandexOAUTHToken) {
		return
	}

	slog.Info("Application doesn't run, starting")

	resultChan := make(chan string)
	wg := sync.WaitGroup{}

	ctx, done := context.WithCancel(context.Background())

	ipcServer := ipc.MustServer(resultChan)
	wg.Add(1)
	go ipc.StartServer(ipcServer, &wg, ctx)
	slog.Info("IPC server started")

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case result := <-resultChan:
				slog.Info("Received IPC result", slog.String("result", result))
			case <-ctx.Done():
				slog.Debug("IPC result handler stopped")
				return
			}
		}
	}()


	p := tea.NewProgram(ui.NewTabs(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		slog.Error("There here is an error", slog.String("Error", err.Error()))
	}

	slog.Debug("Start graceful")
	done()
	wg.Wait()

	slog.Debug("Closing application")
}

func isSendFlagsToClientSuccess(yandexOAUTHToken string) (result bool) {
	result = false
	ipcClient, err := ipc.ConnectToServer()

	if err == nil {
		defer ipc.CloseClient(ipcClient)

		slog.Info("Application already running, setup flags")

		if yandexOAUTHToken != "" {
			err = ipc.SendYandexOAUTHToken(ipcClient, yandexOAUTHToken)
			if err != nil {
				slog.Error("Cannot set yandex oauth token")
			}
		}

		result = true
	}

	return
}
