package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/nifle3/tui_music/internal/ipc"
	"golang.org/x/net/context"
)

var version string

func main() {
	config := mustNewConfig()
	setupLogger(config.LogLevel)

	slog.Debug(fmt.Sprintf("Init program with config %#v", config))
	slog.Info("Logger setup")

	yandexOAUTHToken := flag.String("yandex_oauth", "", "Yandex OAUTH key to access music api via account")
	flag.Parse()

	gracefulCtx, gracefulStop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer gracefulStop()

	if isSendFlagsToClientSuccess(*yandexOAUTHToken) {
		return
	}

	slog.Info("Application doesn't run, starting")

	resultChan := make(chan string)
	wg := sync.WaitGroup{}

	ipcServer := ipc.MustServer(resultChan)
	wg.Add(1)
	go ipc.StartServer(ipcServer, &wg, gracefulCtx)
	defer ipc.StopServer(ipcServer)

	slog.Info("IPC server started")

	wg.Wait()
}

func isSendFlagsToClientSuccess(yandexOAUTHToken string) (result bool) {
	result = false
	ipcClient, err := ipc.ConnectToServer()
	defer ipc.CloseClient(ipcClient)

	if err == nil {
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
