package main

import (
	"testing"

	"log/slog"

	"github.com/stretchr/testify/assert"
)

func TestGetLevel(t *testing.T) {
	t.Parallel()

	data := []struct{
		name string
		input string
		expect slog.Level
	}{
		{
			name: "Correct value DEBUG",
			input: "DEBUG",
			expect: slog.LevelDebug,
		},
		{
			name: "Correct value of WARNING",
			input: "Warning",
			expect: slog.LevelWarn,
		},
		{
			name: "Correct value of ERROR",
			input: "ERROR",
			expect: slog.LevelError,
		},
		{
			name: "Correct value of Info",
			input: "INFO",
			expect: slog.LevelInfo,
		},
		{
			name: "Empty value",
			input: "",
			expect: slog.LevelDebug,
		},
		{
			name: "Spacing value",
			input: "    ",
			expect: slog.LevelDebug,
		},
		{
			name: "Incorect value",
			input: "asdqwe",
			expect: slog.LevelDebug,
		},
		{
			name: "UTF-8",
			input: "🤡🤡🤡🤡",
			expect: slog.LevelDebug,
		},
		{
			name: "Valid prefix",
			input: "INFO🤡🤡🤡🤡",
			expect: slog.LevelDebug,
		},
		{
			name: "Valid sufix",
			input: "🤡🤡🤡INFO",
			expect: slog.LevelDebug,
		},
	}

	for _, value := range data {
		t.Run(value.name, func (t *testing.T) {
			actual := getLevel(value.name)
			assert.Equal(t, value.expect, actual)
		})
	}
}
