package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// Main entry point for the TUI. Initializes the main model
func StartTea() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Printf("Unable to open log file: %v", err)
		os.Exit(1)
	}
	defer f.Close()

	m, _ := InitProject()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Unable to start CLI: %v\n", err)
		os.Exit(1)
	}
}
