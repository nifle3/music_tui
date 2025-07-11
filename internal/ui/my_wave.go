package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type MyWave struct {

}

func NewMyWave() MyWave {
	return MyWave{}
}

func (mw MyWave) Init() tea.Cmd {
	return nil
}

func (mw MyWave) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (mw MyWave) View() string {
	return "MyWave"
}
