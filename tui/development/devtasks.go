package development

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/luka2220/devtasks/database"
)

type devBoardModel struct {
	board    database.Board
	quitting bool
	option   int
}

func StartManageModel() *devBoardModel {
	var err error

	m := &devBoardModel{}
	m.board, err = database.GetActiveBoard()
	m.quitting = false
	if err != nil {
		panic(err)
	}

	return m
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
