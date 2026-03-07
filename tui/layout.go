package tui

import (
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/bubbles/textinput"
)

func renderTopPanels() string {

	history := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(26).
		Height(12).
		Render("History")

	response := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(54).
		Height(12).
		Render("Response")

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		history,
		response,
	)
}

func (m Model) renderRequestBar() string {
	// button style
	sendButton := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 2).
		Render("Send")

		// combine button and input
	content := lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.urlInput.View(),
		" ",
		sendButton,
	)
	requestBarStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(82).Height(3).PaddingLeft(1)
	requestBar := requestBarStyle.Render(content)
	return requestBar
}

func renderTabs() string {

	tab1 := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Render("Body")

	tab2 := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(0, 1).
		Render("Header")

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		tab1,
		tab2,
	)
}

func renderTabContent() string {
	tabContent := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(82).Height(4).PaddingLeft(1).Render("Tab content")
	return tabContent
}
