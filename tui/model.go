package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

// type ActiveTab int

// const (
// 	TabNone ActiveTab = iota
// 	TabBody
// 	TabHeaders
// )

type Model struct {
	width  int
	height int

	focusIndex int

	// tabs state
	MainTab       []string
	MainTabIndex  int
	BodyModeTab   []string
	BodyModeIndex int
	showModeList  bool
	editor        textarea.Model

	LastMessage string

	// request inputs (you'll add real components later)
	methods     []string // current state of method
	methodIndex int
	url         string

	urlInput textinput.Model // input state

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

	editor := textarea.New()
	editor.Placeholder = "Enter json/formData"
	editor.Focus()

	//current state
	return Model{
		urlInput:      ti,
		methods:       []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		methodIndex:   0,
		focusIndex:    0,
		MainTab:       []string{"Body", "Header"},
		MainTabIndex:  0,
		BodyModeTab:   []string{"FormData", "json"},
		BodyModeIndex: 0,
		showModeList:  false,
		editor:        editor,
	}
}
