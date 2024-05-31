package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luka2220/devtasks/constants"
)

func StartProjectInitTui() {
	m := &initializationModel{}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		constants.Logger.WriteString(fmt.Sprintf("Error starting project init tui: %v", err))
		os.Exit(1)
	}
}

type initializationModel struct {
	quitting    bool
	projectName string
}

func (m *initializationModel) Init() tea.Cmd {
	return nil
}

func (m *initializationModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *initializationModel) View() string {
	s := "Project Initialization TUI started 💪"

	if m.quitting {
		s += "\n\nSee you next time! 👋\n"
	}

	return s
}
