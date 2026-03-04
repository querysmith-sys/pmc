package tui

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	// Top panels
	output := m.renderTopPanel()
	return tea.NewView(output)
}
