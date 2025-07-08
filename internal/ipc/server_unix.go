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
	outer: for {
		select {
			case <- ctx.Done():
				break outer
			default:
				conn, err := server.listener.Accept()
				if err != nil {
					slog.Error("Listener accecpt error", slog.String("error", err.Error()))
					continue
				}
				defer conn.Close()

				request, err := handleConnection(conn, ctx)
				if err != nil {
					continue
				}
				server.resultChan <- request
		}
	}
}

func handleConnection(conn net.Conn, _ context.Context) (string, error) {
	conn.SetReadDeadline(time.Now().Add(time.Second * 30))
	data, err := io.ReadAll(conn)
	if err != nil {
		slog.Error("Invalid io.ReadAll error", slog.String("Error", err.Error()))
		return "", err
	}

	return string(data), nil
}

func StopServer(server IpcServer) {
	server.listener.Close()
	os.Remove(socketPath)
}
