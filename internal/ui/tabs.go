package ui

import (
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

var selectedTabStyle = lipgloss.
	NewStyle().
	Bold(true).
	Background(lipgloss.CompleteColor{
        TrueColor: "#585858",
        ANSI256:   "240",
        ANSI:      "8",
    }).
	Inherit(tabStyle)

var tabPanelStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.CompleteColor{
        TrueColor: "#585858",
        ANSI256:   "240",
        ANSI:      "8",
    }).
	BorderLeft(false).BorderTop(false).BorderRight(true).BorderBottom(false)

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

	height int
}

func NewTabs() Tabs {
	return Tabs{
		tabs: []Tab{
			{
				name:  "Моя волна",
				model: NewMyWave(),
			},
			{
				name:  "Рекомендации",
				model: NewRecomendation(),
			},
			{
				name:  "Мои плейлисты",
				model: NewMyPlaylists(),
			},
			{
				name:  "Лайки",
				model: NewLikes(),
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
	case tea.WindowSizeMsg:
		t.height = msg.Height
		return t, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if t.cursor+1 > len(t.tabs)-1 {
				t.cursor = 0
			} else {
				t.cursor++
			}

			return t, nil

		case "1":
			t.cursor = 0
			return t, nil

		case "2":
			t.cursor = 1
			return t, nil

		case "3":
			t.cursor = 2
			return t, nil

		case "4":
			t.cursor = 3
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

	listOfTabs := builder.String()
	tabPanel := tabPanelStyle.Height(t.height).Render(listOfTabs)

	selectedTab := t.tabs[t.cursor]
	selectedTabView := selectedTab.model.View()

	view := lipgloss.JoinHorizontal(lipgloss.Top, tabPanel, selectedTabView)

	return view
}
