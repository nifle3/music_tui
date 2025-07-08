package ipc

import "net/http"

type Client http.Client

func NewClient() Client {
	return Client(http.Client{})
}

func (c Client) IsServerAvailable() (bool, error) {
	return true, nil
}

func (c Client) SetYandexOAUTHToken(token string) error {
	return nil
}
