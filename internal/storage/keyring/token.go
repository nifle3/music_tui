package secrets

import (
	"github.com/nifle3/tui_music/internal/storage"
	"github.com/zalando/go-keyring"
)

var _ storage.Token = Storage{}

type Storage struct {
	appName string
	keyName string
	userName string
}

func New(appName string, userName string) Storage {
	return Storage{
		appName: appName,
		userName: userName,
	}
}

func (s Storage) WithAppName(appName string) Storage {
	s.appName = appName
	return s
}

func (s Storage) Get() (string, error) {
	return keyring.Get(s.appName, s.userName)
}

func (s Storage) Set(token string) error {
	return keyring.Set(s.appName, s.userName, token)
}
