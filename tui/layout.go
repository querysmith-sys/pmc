package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/bubbles/textinput"
)

var normalBox = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder())

var focusedBox = lipgloss.NewStyle().
	Border(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("205"))

func (m Model) renderTopPanels() string {
	historyStyle := normalBox
	if m.focusIndex == 5 {
		historyStyle = focusedBox
	}
	history := historyStyle.
		Width(26).
		Height(12).
		Render("History")

	responseStyle := normalBox
	if m.focusIndex == 6 {
		responseStyle = focusedBox
	}
	response := responseStyle.
		Width(54).
		Height(12).
		Render("Response")

	topPanels := lipgloss.JoinHorizontal(
		lipgloss.Top,
		history,
		response,
	)
	return topPanels
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

func (m Model) renderTabs() string {
	tab1Style := normalBox
	if m.focusIndex == 3 {
		tab1Style = focusedBox
	}
	tab1 := tab1Style.
		Padding(0, 1).
		Render("Body")

	tab2Style := normalBox
	if m.focusIndex == 4 {
		tab2Style = focusedBox
	}

	tab2 := tab2Style.
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

func (m Model) renderBodyModeList() string {
	style := lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(lipgloss.Color("#00FFFF")).
		Bold(true)

	selectedStyle := lipgloss.NewStyle().
		PaddingLeft(2).
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true)

	var items strings.Builder

	if m.showModeList {
		for i, item := range m.BodyModeTab {
			if i == m.BodyModeIndex {
				items.WriteString(selectedStyle.Render("> "+item) + "\n")
			} else {
				items.WriteString(style.Render("  "+item) + "\n")
			}
		}
	}

	return items.String()
}
