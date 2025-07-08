package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/nifle3/yandex_music_tui/internal/ipc"
	)

var version string

func main() {
	config := mustNewConfig()
	setupLogger(config.LogLevel)

	slog.Debug(fmt.Sprintf("Init program with config %#v", config))
	slog.Info("Logger setup")

	yandexOAUTHToken := flag.String("yandex_oauth", "", "Yandex OAUTH key to access music api via account")
	flag.Parse()

	fmt.Println("yandex OAUTH token ", *yandexOAUTHToken)

	ipcClient := ipc.NewClient()
	isAlreadyApplicationRun, err := ipcClient.IsServerAvailable()
	if err != nil {
		slog.Error("Invalid ipc server on port")
		os.Exit(0)
	}

	if isAlreadyApplicationRun {
		slog.Info("Application already running, setup flags")

		if *yandexOAUTHToken != "" {
			err = ipcClient.SetYandexOAUTHToken(*yandexOAUTHToken)
			if err != nil {
				slog.Error("Cannot set yandex oauth token")
			}
		}

		os.Exit(1)
	}

	slog.Info("Application doesn't run, starting")

	go ipc.StartServer()

	slog.Info("IPC server started")
}
