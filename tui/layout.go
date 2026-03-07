package tui

import (
	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/bubbles/textinput"
)

var normalBox = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder())

var focusedBox = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("205"))

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
	sendStyle := normalBox
	if m.focusIndex == 2 {
		sendStyle = focusedBox
	}
	sendButton := sendStyle.
		Padding(0, 1).
		Render("Send")

	// method stylr
	methodStyle := normalBox
	if m.focusIndex == 0 {
		methodStyle = focusedBox
	}

	urlBox := normalBox.Render(m.urlInput.View())

	if m.focusIndex == 1 {
		urlBox = focusedBox.Render(m.urlInput.View())
	}

	methodBox := methodStyle.
		Width(6).
		Render(m.methods[m.methodIndex])

	// input style

	if m.focusIndex == 1 {
		m.urlInput.Focus()
	} else {
		m.urlInput.Blur()
	}
	// combine button and input and method
	content := lipgloss.JoinHorizontal(
		lipgloss.Left,
		methodBox,
		"",
		urlBox,
		"",
		sendButton,
	)
	requestBarStyle := lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Width(82).PaddingLeft(1)
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
