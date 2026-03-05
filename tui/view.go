package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() tea.View {

	if m.width < 60 || m.height < 20 {
		return tea.NewView("Terminal too small")
	}

	top := renderTopPanels()
	requestBar := renderRequestBar()
	tabs := renderTabs()
	tabContent := renderTabContent()

	// combine layout
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		top,
		requestBar,
		tabs,
		tabContent,
	)

	// center it inside terminal
	centered := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		content,
	)

	return tea.NewView(centered)
}
