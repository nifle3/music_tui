package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"fmt"
)

type track struct {
	author string
	title string
	duration string
}

func (t track) Title() string       { return t.title }
func (t track) Description() string { return fmt.Sprintf("%s • %s", t.author, t.duration) }
func (t track) FilterValue() string { return t.title }

type Likes struct {
	list list.Model
}

func NewLikes() Likes {
	items := []list.Item{
		track{
			author: "QWE",
			title: "Qwe",
			duration: "Qwe",
		},
		track{
			author: "QWE",
			title: "Qwe",
			duration: "Qwe",
		},
		track{
			author: "QWE",
			title: "Qwe",
			duration: "Qwe",
		},
		track{
			author: "QWE",
			title: "Qwe",
			duration: "Qwe",
		},
		track{
			author: "QWE",
			title: "Qwe",
			duration: "Qwe",
		},
		track{
			author: "QWE",
			title: "Qwe",
			duration: "Qwe",
		},
		track{
			author: "QWE",
			title: "Qwe",
			duration: "Qwe",
		},
		track{
			author: "QWE",
			title: "Qwe",
			duration: "Qwe",
		},
	}
	list := list.New(items, list.NewDefaultDelegate(), 0, 0)
	list.Title = "Мои лайки"
	return Likes{list: list}
}

func (l Likes) Init() tea.Cmd {
	return nil
}

func (l Likes) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			l.list.SetSize(msg.Width, msg.Height)
		}

	var cmd  tea.Cmd
	l.list, cmd = l.list.Update(msg)
	return l, cmd
}

func (l Likes) View() string {
	return l.list.View()
}
