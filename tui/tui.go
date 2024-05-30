package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luka2220/dev-tracker/constants"
)

// Main entry point for the TUI. Initializes the main model
func StartTea() {
	var err error

	constants.Logger, err = tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Printf("Unable to open log file: %v", err)
		os.Exit(1)
	}
	defer constants.Logger.Close()

	StartMenu()
}
