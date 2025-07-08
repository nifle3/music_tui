//go:build (linux || darwin || freebsd)

package ipc

const (
	socketPath = "/tmp/tui_music.socket"
)
