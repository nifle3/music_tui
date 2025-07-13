package keyring

import (
	"errors"
	"os/user"

	"github.com/zalando/go-keyring"
)

func IsAvailable(appName string, user *user.User) bool {
	if err := keyring.Set(appName, user.Name, "test"); err != nil {
		if errors.Is(err, keyring.ErrUnsupportedPlatform) {
			return false
		}
	}

	if err := keyring.Delete(appName, user.Name); err != nil {
		if errors.Is(err, keyring.ErrUnsupportedPlatform) {
			return false
		}
	}

	return true
}
