package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	// tea "charm.land/bubbletea/v2"
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

	urlInput textinput.Model
	choices  []string // e.g., []string{"Option 1", "Option 2"}
	cursor   int      // which item the user is pointing at
	selected string   // the final selected item

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

	// current state of the url input
	ti := textinput.New()
	ti.Placeholder = "Enter URL..."
	ti.Width = 60
	ti.Focus()

	return Model{
		urlInput: ti,
	}
	// return Model{
	// 	method:    "GET",
	// 	activeTab: TabNone,
	// }
}
