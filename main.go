package main

import (
	"flag"
	"pmc/cmd"
	"pmc/tui"

	tea "github.com/charmbracelet/bubbletea"
)

var uimode = flag.Bool("ui", false, "Launch the application in TUI mode")

func main() {
	flag.Parse()
	if *uimode {
		p := tea.NewProgram(tui.InitModel())
		if _, err := p.Run(); err != nil {
			panic(err)
		}
		return // <-- VERY IMPORTANT
	}
	cmd.ExecuteCLI()
}
