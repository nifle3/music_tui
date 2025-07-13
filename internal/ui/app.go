package ui

import (
	"log/slog"
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	loginPage tea.Model
	yandexPage tea.Model
	helpPage tea.Model

	helpPageToggle bool
	isUserAuth bool
}

func NewApp() App {
	return App{
		helpPageToggle: false,
		isUserAuth: false,
		loginPage: NewLogin(),
	}
}

func (a App) Init() tea.Cmd {
	if a.isUserAuth {
		return a.yandexPage.Init()
	}

	return a.loginPage.Init()
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case LoginMsg:
			a.isUserAuth = true
			return a, a.yandexPage.Init()

		case LogoutMsg:
			a.isUserAuth = false
			return a, a.loginPage.Init()

		case tea.KeyMsg:
			switch msg.String() {
				case "h", "р":
					a.helpPageToggle = !a.helpPageToggle
					return a, nil

				case "ctrl+c", "ctrl+с", "esc":
					slog.Debug("Closing application")
					return a, tea.Quit
			}
	}

	var cmd tea.Cmd
	if a.isUserAuth {
		a.yandexPage, cmd = a.yandexPage.Update(msg)
	} else {
		a.loginPage, cmd = a.loginPage.Update(msg)
	}

	return a, cmd
}

func (a App) View() string {
	if a.helpPageToggle {
		return a.helpPage.View()
	}

	if a.isUserAuth {
		return a.yandexPage.View()
	}

	return a.loginPage.View()
}
