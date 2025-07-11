package ui

import tea "github.com/charmbracelet/bubbletea"

type MyPlaylists struct {

}

func NewMyPlaylists() MyPlaylists {
	return MyPlaylists{}
}

func (ml MyPlaylists) Init() tea.Cmd {
	return nil
}

func (ml MyPlaylists) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return ml, nil
}

func (ml MyPlaylists) View() string {
	return "Playlists"
}
