package tui

import "github.com/charmbracelet/lipgloss"





func (m Model) renderTopPanel() string {
	topPanel := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Padding(1).Render("hello")
	return topPanel
}

func renderRequestbar() {

}

func renderTabs() {

}

func renderTabContent() {

}
