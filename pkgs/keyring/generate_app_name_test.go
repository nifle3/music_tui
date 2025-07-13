package keyring

import (
	"fmt"
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAppName(t *testing.T) {
	t.Parallel()

	data := []struct{
		expected string
		inputUser user.User
		inputAppName string
	}{
		{
			expected: "cd89a17839ec88ead08af743265767cabb9d46d143fd4f91c32172d7c2984966",
			inputUser: user.User{
				Username: "qwe",
				Uid: "qwe",
				Gid: "qwe",
				Name: "qwe",
				HomeDir: "/home/dir",
			},
			inputAppName: "appname",
		},
		{
			expected: "96dcc8656aac72548c3473a1bd53571e971ebcf9790561a3bfab6094c97f51fa",
			inputUser: user.User{
				Username: "qwe",
				Uid: "9999",
				Gid: "qwe",
				Name: "qwe",
				HomeDir: "/home/dir",
			},
			inputAppName: "appname",
		},
		{
			expected: "b4f34187cbd99bc53769a019171cc88f95fc5572fe40d82d77936383094f658a",
			inputUser: user.User{
				Username: "qwe",
				Uid: "1001",
				Gid: "qwe",
				Name: "qwe",
				HomeDir: "/home/dir",
			},
			inputAppName: "appname",
		},
		{
			expected: "ab80693b6e61b20613fff69ffb3458ce2f6834f11b2fb009b31d2ff3154d8d24",
			inputUser: user.User{
				Username: "qwe",
				Uid: "0000000000000000000000000000000000000000000000000000",
				Gid: "qwe",
				Name: "qwe",
				HomeDir: "/home/dir",
			},
			inputAppName: "appname",
		},
	}

	for idx, value := range data {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			result := GenerateAppName(value.inputAppName, &value.inputUser)

			assert.Equal(t, value.expected, result)
		})
	}
}
