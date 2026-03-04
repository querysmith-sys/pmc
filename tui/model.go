package tui

import (
	"fmt"
	"time"

	tea "charm.land/bubbletea/v2"
)

type ActiveTab int

const (
	TabNone ActiveTab = iota
	TabBody
	TabHeaders
)

type Model struct {
	width  int
	height int

	focusIndex int
	activeTab  ActiveTab

	LastMessage string

	// request inputs (you'll add real components later)
	method string
	url    string

	bodyInput    string
	headersInput string

	// response display
	response   string
	statusCode int
	timeTaken  time.Duration

	// history
	history []HistoryEntry

	loading bool
}

type HistoryEntry struct {
	Method     string
	URL        string
	StatusCode int
	TimeTaken  time.Duration
}

// type Model struct {
// 	body           string
// 	status         string
// 	headerDetails  map[string]string
// 	responsemethod string
// 	contentLength  int
// 	statusCode     int
// 	timeTaken      time.Duration
// 	History        []HistoryEntry
// 	BodyContent    []byte
// 	HeadeContent   []byte
// }

func InitModel() Model {
	return Model{
		method:    "GET",
		activeTab: TabNone,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() tea.View {
	content := "Hello, World!"
	v := content
	r := fmt.Sprintf("Width: %d\nHeight: %d\n\n %s\n\nPress q to quit.", m.width, m.height, v)
	return tea.NewView(r)
	// result1 := tea.NewView(v)
	// result2 := tea.NewView(r)
	// return  result1, result2
}
