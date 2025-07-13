package keyring

import (
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zalando/go-keyring"
)

func TestIsAvailable_NotAvailable(t *testing.T) {
	keyring.MockInitWithError(keyring.ErrUnsupportedPlatform)
	result := IsAvailable("qwe", &user.User{
		Name: "Qwe",
	})

	assert.False(t, result)
}

func TestIsAvailable_AvailableWithAnotherError(t *testing.T) {
	keyring.MockInitWithError(keyring.ErrSetDataTooBig)
	result := IsAvailable("qwe", &user.User{
		Name: "Qwe",
	})

	assert.True(t, result)
}

func TestIsAvailable_Available(t *testing.T) {
	keyring.MockInit()
	result := IsAvailable("qwe", &user.User{
		Name: "Qwe",
	})

	assert.True(t, result)
}
