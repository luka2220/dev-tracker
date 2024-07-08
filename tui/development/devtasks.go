package development

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luka2220/devtasks/constants"
	"github.com/luka2220/devtasks/database"
)

// Start the tui app associated with the currently active board
func StartDevTaskBoard() {
	var err error
	m := &devBoardModel{}
	m.board, err = database.GetCurrentActiveBoard()
	m.quitting = false
	if err != nil {
		panic(err)
	}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		constants.Logger.WriteString(fmt.Sprintf("Error starting dev task board tui: %v", err))
		os.Exit(1)
	}
}

type devBoardModel struct {
	board    database.Board
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
	s := fmt.Sprintf("Active Board: %s", m.board.Name)

	if m.quitting {
		s += "\n\nSee you next time! 👋\n"
	}

	return s
}
