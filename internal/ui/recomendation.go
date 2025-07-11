package ui

import tea "github.com/charmbracelet/bubbletea"

type Recomendation struct {

}

func NewRecomendation() Recomendation {
	return Recomendation{}
}

func (r Recomendation) Init() tea.Cmd {
	return nil
}

func (r Recomendation) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return r, nil
}

func (r Recomendation) View() string {
	return "Recomendation"
}
