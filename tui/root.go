package tui

import tea "github.com/charmbracelet/bubbletea"

type activeModelIdx int

const (
	createBoardModelIdx activeModelIdx = iota
	manageBoardModelIdx
)

type rootModel struct {
	quitting    bool
	activeModel activeModelIdx
}

func (m *rootModel) Init() tea.Cmd {
	return nil
}

func (m *rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *rootModel) View() string { return "" }
