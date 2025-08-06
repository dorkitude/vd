package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MainMenuModel struct {
	choices  []string
	cursor   int
	selected string
	width    int
	height   int
}

func NewMainMenuModel() *MainMenuModel {
	return &MainMenuModel{
		choices: []string{"📚 Browse Documentation", "➕ Add New Project", "🚪 Exit"},
	}
}

func (m *MainMenuModel) Init() tea.Cmd {
	return nil
}

func (m *MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = m.choices[m.cursor]
			switch m.cursor {
			case 0:
				return NewBrowseModel(), nil
			case 1:
				return NewAddModel(), nil
			case 2:
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, nil
}

func (m *MainMenuModel) View() string {
	s := TitleStyle.Render("🚀 Various Docs")
	s += "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		style := ItemStyle
		if m.cursor == i {
			cursor = "▸"
			style = SelectedItemStyle
		}

		s += fmt.Sprintf("%s %s\n", cursor, style.Render(choice))
	}

	s += "\n"
	s += HelpStyle.Render("↑/↓ or j/k to navigate • enter to select • q to quit")

	centeredStyle := lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center)

	if m.width > 0 && m.height > 0 {
		return centeredStyle.Render(s)
	}

	return DocStyle.Render(s)
}
