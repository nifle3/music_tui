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
			input: "Warn",
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
			expect: slog.LevelInfo,
		},
		{
			name: "Spacing value",
			input: "    ",
			expect: slog.LevelInfo,
		},
		{
			name: "Incorect value",
			input: "asdqwe",
			expect: slog.LevelInfo,
		},
		{
			name: "UTF-8",
			input: "🤡🤡🤡🤡",
			expect: slog.LevelInfo,
		},
		{
			name: "Valid prefix",
			input: "DEBUG🤡🤡🤡🤡",
			expect: slog.LevelInfo,
		},
		{
			name: "Valid sufix",
			input: "🤡🤡🤡DEBUG",
			expect: slog.LevelInfo,
		},
	}

	for _, value := range data {
		t.Run(value.name, func (t *testing.T) {
			actual := getLevel(value.input)
			assert.Equal(t, value.expect, actual)
		})
	}
}
