package ui

import tea "github.com/charmbracelet/bubbletea"

type Likes struct {

}

func NewLikes() Likes {
	return Likes{}
}

func (l Likes) Init() tea.Cmd {
	return nil
}

func (l Likes) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return l, nil
}

func (l Likes) View() string {
	return "Likes"
}
