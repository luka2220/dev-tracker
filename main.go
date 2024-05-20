package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luka2220/dev-tracker/constants"
	"github.com/luka2220/dev-tracker/tui"
)

func main() {
	m, _ := tui.InititProject()
	constants.P = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.P.Run(); err != nil {
		fmt.Printf("Unable to start CLI: %v\n", err)
		os.Exit(1)
	}
}
