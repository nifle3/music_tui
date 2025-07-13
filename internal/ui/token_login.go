package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type TokenLogin struct {
	input textinput.Model
	err error
}

func NewTokenLogin() TokenLogin {
	textField := textinput.New()
	textField.Placeholder = "Yandex music token"
	textField.Focus()
	textField.Width = 20

	return TokenLogin{
		input: textField,
	}
}

func (t TokenLogin) Init() tea.Cmd {
	return textinput.Blink
}

func (t TokenLogin) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	t.input, cmd = t.input.Update(msg)
	return t, cmd
}

func (t TokenLogin) View() string {
	return t.input.View()
}
