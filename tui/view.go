package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {

	if m.width < 60 || m.height < 20 {
		return "Terminal too small"
	}

	top := m.renderTopPanels()
	requestBar := m.renderRequestBar()
	tabs := m.renderTabs()
	tabContent := renderTabContent()
	list := m.renderBodyModeList()
	// combine layout
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		top,
		requestBar,
		tabs,
		list,
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

	return centered
}
