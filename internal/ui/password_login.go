package ui

import tea "github.com/charmbracelet/bubbletea"

type PasswordLogin struct {

}

func NewPasswordLogin() PasswordLogin {
	return PasswordLogin{}
}

func (p PasswordLogin) Init() tea.Cmd {
	return nil
}

func (p PasswordLogin) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (p PasswordLogin) View() string {
	return "Password login"
}
