//go:build linux || darwin || freebsd

package ipc

import (
	"context"
	"io"
	"log/slog"
	"net"
	"os"
	"sync"
	"time"
)

type IpcServer struct {
	resultChan chan<- string
	listener net.Listener
}

func MustServer(resultChan chan<- string) IpcServer {
	socket, err := net.Listen("unix", socketPath)
	if err != nil {
		slog.Error("cannot create listener")
		panic(err.Error())
	}

	return IpcServer {
		resultChan: resultChan,
		listener: socket,
	}
}

func StartServer(server IpcServer, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	go func() {
		<-ctx.Done()
		slog.Debug("Context done, closing listener")
		server.listener.Close()
		os.Remove(socketPath)
	}()

	for {
		conn, err := server.listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				slog.Debug("Listener closed due to context done")
				return
			default:
				slog.Error("Listener accept error", slog.String("error", err.Error()))
				continue
			}
		}

		request, err := handleConnection(conn, ctx)
		if err != nil {
			slog.Error("Handle connection error", slog.String("error", err.Error()))
			return
		}
		server.resultChan <- request
	}
}

func handleConnection(conn net.Conn, _ context.Context) (string, error) {
	defer conn.Close()

	slog.Debug("Start handle connection")
	conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	data, err := io.ReadAll(conn)
	if err != nil {
		slog.Error("Invalid io.ReadAll error", slog.String("Error", err.Error()))
		return "", err
	}

	slog.Debug("End handle connection")
	return string(data), nil
}
