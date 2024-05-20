package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
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
func InititProject() (tea.Model, tea.Cmd) {
	width, height := 20, 14
	items := []list.Item{
		item("Add"),
		item("View"),
		item("Edit"),
	}

	// Create the list
	l := list.New(items, itemDelegate{}, width, height)
	l.Title = "Select an operation"

	// Defining a text input component
	ti := textinput.New()
	// ti.Focus()
	ti.Prompt = "> "
	ti.Placeholder = "Prompt..."
	ti.CharLimit = 250
	ti.Width = 50

	m := Model{session: nav, input: ti, options: l}

	return m, nil
}

type mode int

type Model struct {
	session  mode
	input    textinput.Model
	options  list.Model
	selected string
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
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.options, cmd = m.options.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	s := "Task Operations\n"

	var titleStyle = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Foreground(mainFGC)

	cmd := "To quit: ctrl+c/q"
	var commandsStyle = lipgloss.NewStyle().
		SetString(cmd).
		Bold(true).
		Italic(true).
		Foreground(mainFGC)

	var promptStyle = lipgloss.NewStyle().
		SetString(m.input.View()).
		Bold(true).
		Foreground(promptFGC)

	titleRendered := titleStyle.Render()
	commandRendered := commandsStyle.Render() + "\n"
	promptRendered := promptStyle.Render()

	if m.input.Focused() {
		return lipgloss.JoinVertical(lipgloss.Top, titleRendered, commandRendered, promptRendered)
	}

	return m.options.View()
}
