package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab":
			m.focusIndex = (m.focusIndex + 1) % 3
		case "shift+tab":
			m.focusIndex = (m.focusIndex + 2) % 3
		case "right":
			if m.focusIndex == 0 {
				m.methodIndex = (m.methodIndex + 1) % len(m.methods)
			}

		case "left":
			if m.focusIndex == 0 {
				m.methodIndex = (m.methodIndex - 1 + len(m.methods)) % len(m.methods)
			}

		}
		if m.focusIndex == 1 {
			m.urlInput, cmd = m.urlInput.Update(msg)
		}
	}

	return m, cmd
}
