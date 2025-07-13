package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type LoginMethod struct {
	model tea.Model
	name string
}

type Login struct {
	loginMethods []LoginMethod
	cursor int
}

func NewLogin() Login {
	return Login{
		loginMethods: []LoginMethod{
			{
				model: NewPasswordLogin(),
				name: "По паролю",
			},
			{
				model: NewTokenLogin(),
				name: "По токену",
			},
		},
	}
}

func (l Login) Init() tea.Cmd {
	return l.loginMethods[l.cursor].model.Init()
}

func (l Login) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if l.cursor+1 > len(l.loginMethods)-1 {
				l.cursor = 0
			} else {
				l.cursor++
			}

			return l, l.loginMethods[l.cursor].model.Init()
		}
	}

	return l, nil
}

func (l Login) View() string {
	return l.loginMethods[l.cursor].model.View()
}
