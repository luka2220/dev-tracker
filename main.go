package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/luka2220/dev-tracker/constants"
)

func main() {
	var err error

	constants.Logger, err = tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Printf("Unable to open log file: %v", err)
		os.Exit(1)
	}
	defer constants.Logger.Close()

	m := RootModel{}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		constants.Logger.WriteString(fmt.Sprintf("Error Starting TUI: %v", err))
		os.Exit(1)
	}
}

// NOTE:
// Root/Base model for the programs structure
// menuInput: field for adding a bubble textInput model to the TUI
// operation: menu operation code inputted by the user
type RootModel struct {
	menuInput textinput.Model
	operation int
}

func (m RootModel) Init() tea.Cmd {
	return nil
}

func (m RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m RootModel) View() string {
	return "Prog running"
}
