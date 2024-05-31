package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luka2220/devtasks/constants"
)

func StartDevTaskBoard() {
	m := &devBoardModel{}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		constants.Logger.WriteString(fmt.Sprintf("Error starting dev task board tui: %v", err))
		os.Exit(1)
	}
}

type devBoardModel struct {
	quitting bool
	option   int
}

func (m *devBoardModel) Init() tea.Cmd {
	return nil
}

func (m *devBoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m *devBoardModel) View() string {
	s := "Development Task Board TUI started 💪"

	if m.quitting {
		s += "\n\nSee you next time! 👋\n"
	}

	return s
}
