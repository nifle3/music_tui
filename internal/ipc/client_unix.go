//go:build linux || darwin || freebsd

package ipc

import (
	"net"
	"time"
)

type Client net.Conn

func ConnectToServer() (Client, error) {
	conn, err := net.Dial("unix", socketPath)
	return Client(conn), err
}

func SendYandexOAUTHToken(client Client, token string) error {
	client.SetWriteDeadline(time.Now().Add(time.Second * 20))

	_, err := client.Write([]byte(token))

	return err
}

func CloseClient(client Client) {
	client.Close()
}
