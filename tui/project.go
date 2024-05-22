package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
type item string

func (i item) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int  { return 1 }
func (d itemDelegate) Spacing() int { return 0 }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
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
func InitProject() (tea.Model, tea.Cmd) {
	width, height := 20, 14
	items := []list.Item{
		item("Add"),
		item("View"),
		item("Edit"),
	}

	// Create the list
	l := list.New(items, itemDelegate{}, width, height)
	l.Title = "Select an operation"

	m := Model{session: nav, options: l}

	return m, nil
}

type mode int

type Model struct {
	session  mode
	options  list.Model
	selected string
	chosen   bool
	quitting bool
}

func initModel() *Model {
	return &Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	}

	if m.chosen {
		tableUpdateLoop(msg, m)
	}

	return menuUpdateLoop(msg, m)
}

func (m Model) View() string {
	var s string

	if m.chosen {
		s = tableView(m)
	} else {
		s = menuView(m)
	}

	if m.quitting {
		s = "\nSee you next time!!\n\n"
	}

	return s
}

// Main menu view update loop
func menuUpdateLoop(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			i, ok := m.options.SelectedItem().(item)
			if ok {
				m.selected = string(i)
				m.chosen = true
			}
		}
	}

	var cmd tea.Cmd
	m.options, cmd = m.options.Update(msg)

	return m, cmd
}

// Table view update loop
func tableUpdateLoop(_ tea.Msg, m Model) (tea.Model, tea.Cmd) {
	return m, nil
}

func menuView(m Model) string {
	return m.options.View()
}

func tableView(_ Model) string {
	s := "\nInside the table view!!!!\n\n"
	return s
}
