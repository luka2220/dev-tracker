package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/luka2220/devtasks/constants"
	"github.com/luka2220/devtasks/tui/development"
	"github.com/luka2220/devtasks/tui/initialization"
)

const (
	rootBoardModelIdx activeModelIdx = iota
	initBoardModelIdx
	manageBoardModelIdx
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type activeModelIdx int

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func InitializeRootTui() *rootModel {
	items := []list.Item{
		item{title: "Initialize", desc: "initalize a new board for development tracking"},
		item{title: "Manage", desc: "manage any existing boards"},
	}
	menuListModel := list.New(items, list.NewDefaultDelegate(), 0, 0)

	m := &rootModel{quitting: false, activeModel: rootBoardModelIdx, menuList: menuListModel}
	m.menuList.Title = "Select an option 🤖"

	return m
}

func StartRootTui() {
	m := InitializeRootTui()

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		constants.Logger.WriteString(fmt.Sprintf("Error starting dev task board tui: %v", err))
		os.Exit(1)
	}
}

type rootModel struct {
	active      tea.Model
	quitting    bool
	activeModel activeModelIdx
	menuList    list.Model
}

func (m *rootModel) Init() tea.Cmd {
	return nil
}

func (m *rootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			v := m.menuList.SelectedItem().FilterValue()
			m.setActiveModel(v)
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.menuList.SetSize(msg.Width-h, msg.Height-v)
	}

	if m.activeModel == initBoardModelIdx {
		// Start the init board bubble tea model
		return initialization.StartInitModel(), cmd
	}

	if m.activeModel == manageBoardModelIdx {
		// start the manage board bubble tea model
		return development.StartManageModel(), cmd
	}

	m.menuList, cmd = m.menuList.Update(msg)
	return m, cmd
}

func (m *rootModel) View() string {
	s := "Welcome to the development tracker CLI 🤩\n\n"

	s += fmt.Sprintf("%s\n", docStyle.Render(m.menuList.View()))

	if m.quitting {
		s += "See you next time 👋"
	}

	return s
}

// Updates the current bubble tea model to show different views
func (m *rootModel) setActiveModel(selection string) {
	switch selection {
	case "Initialize":
		m.activeModel = initBoardModelIdx
	case "Manage":
		m.activeModel = manageBoardModelIdx
	}
}
