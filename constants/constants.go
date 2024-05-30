package constants

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// NOTE:
/* PROGRAM CONSTANTS */
var (
	P      *tea.Program // Current bubble tea program running
	Logger *os.File     // Global file logger
)
