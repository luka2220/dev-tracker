package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func StartTableView() {
	m := initTableModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Unable to start Table View: %v", err)
		os.Exit(1)
	}
}

type TableModel struct {
	selected     string
	quitting     bool
	previousView bool
}

func initTableModel() *TableModel {
	return &TableModel{}
}

func (m TableModel) Init() tea.Cmd {
	return nil
}

func (m TableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc":
			m.previousView = true
		}
	}

	if m.previousView {
		StartMenu()
	}

	return m, nil
}

func (m TableModel) View() string {
	return "\nInside the Table View Program\n\n"
}
