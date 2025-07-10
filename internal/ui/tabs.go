package ui

import (
	"log/slog"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var tabStyle = lipgloss.
	NewStyle().
	Foreground(lipgloss.CompleteColor{
        TrueColor: "#FFFFFF",
        ANSI256:   "15",
        ANSI:      "7",
    })

var selectedTabStyle = tabStyle.
	Bold(true).
	Background(lipgloss.CompleteColor{
        TrueColor: "#585858",
        ANSI256:   "240",
        ANSI:      "8",
    })

type Tab struct {
	name  string
	model tea.Model
}

func NewTab(name string, model tea.Model) Tab {
	return Tab{
		name:  name,
		model: model,
	}
}

type Tabs struct {
	tabs   []Tab
	cursor int
}

func NewTabs() Tabs {
	return Tabs{
		tabs: []Tab{
			{
				name:  "Моя волна",
				model: nil,
			},
			{
				name:  "Рекомендации",
				model: nil,
			},
			{
				name:  "Мои плейлисты",
				model: nil,
			},
			{
				name:  "Лайки",
				model: nil,
			},
		},
		cursor: 0,
	}
}

func (t Tabs) Init() tea.Cmd {
	return nil
}

func (t Tabs) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			slog.Debug("Closing application")
			return t, tea.Quit

		case "tab":
			if t.cursor+1 > len(t.tabs)-1 {
				t.cursor = 0
			} else {
				t.cursor++
			}

			return t, nil

		case "shift+tab":
			if t.cursor-1 < 0 {
				t.cursor = len(t.tabs)-1
			} else {
				t.cursor--
			}

			return t, nil
		}
	}

	return t, nil
}

func (t Tabs) View() string {
	builder := strings.Builder{}

	for idx, value := range t.tabs {
		if t.cursor == idx {
			builder.WriteString(selectedTabStyle.Render(value.name))
		} else {
			builder.WriteString(tabStyle.Render(value.name))
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
