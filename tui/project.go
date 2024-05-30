package tui

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/luka2220/dev-tracker/constants"
)

/* STYLES */
var (
	mainFGC           = lipgloss.Color("#DFD0B8")
	promptFGC         = lipgloss.Color("#F5EFE6")
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
)

const (
	nav = iota
	edit
	create
)

// NOTE:
// The value that is being filtered when filtering though the list
type MenuItem string

func (i MenuItem) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int  { return 1 }
func (d itemDelegate) Spacing() int { return 0 }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(MenuItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := lipgloss.NewStyle().Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

// NOTE:
// Initializes the based project tui model
func StartMenu() {
	width, height := 20, 14
	items := []list.Item{
		MenuItem("Add"),
		MenuItem("View"),
		MenuItem("Edit"),
	}

	// Create the list
	l := list.New(items, itemDelegate{}, width, height)
	l.Title = "Select an operation"

	m := MainModel{session: nav, Options: l}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		constants.Logger.WriteString(fmt.Sprintf("Unable to start CLI: %v", err))
		os.Exit(1)
	}
}

type mode int

type MainModel struct {
	session  mode
	Options  list.Model
	Selected string
	Chosen   bool
	quitting bool
}

func initModel() *MainModel {
	return &MainModel{}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

// NOTE:
// Main update loop
// Executes a different update listerner depending on the state of the model
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}

	if m.Chosen {
		StartTableView()
	}

	return menuUpdateLoop(msg, m)
}

// NOTE:
// Main Bubbletea View
// Renders different subviews based on the models state
func (m MainModel) View() string {
	var s string

	if m.quitting {
		return "\nSee you next time!!\n\n"
	}

	s = m.Options.View()

	return s
}

// NOTE:
// Main menu sub-view update loop
func menuUpdateLoop(msg tea.Msg, m MainModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i, ok := m.Options.SelectedItem().(MenuItem)
			if ok {
				m.Selected = string(i)
				m.Chosen = true
			}
		}
	}

	var cmd tea.Cmd
	m.Options, cmd = m.Options.Update(msg)

	return m, cmd
}
