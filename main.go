package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/* STYLES */
var (
	bgc               = lipgloss.Color("#153448")
	fgc               = lipgloss.Color("#DFD0B8")
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

const (
	defaultWidth = 50
	listHeight   = 10
)

type item string

func (i item) FilterValue() string { return "" }

// NOTE:
// Describes the general functionality for the list items
type itemDelegate struct{}

// NOTE:
// Redners the items view
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

// NOTE:
// Sets the height of the list item
func (d itemDelegate) Height() int { return 1 }

// NOTE:
// Sets the size of the horizontal gap between list items
func (d itemDelegate) Spacing() int { return 0 }

// NOTE:
// Update loops for the list (similar to update for tea.Model) all messages dome through here
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func main() {
	items := []list.Item{
		item("Login"),
		item("Sign Up"),
	}

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Login or create an account to continue"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	m := MainModel{options: l}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Unable to start CLI: %v\n", err)
		os.Exit(1)
	}
}

type MainModel struct {
	options list.Model
	choice  string
	width   int
	height  int
}

func initModel() *MainModel {
	return &MainModel{}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			i, ok := m.options.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit

		}
	}

	var cmd tea.Cmd
	m.options, cmd = m.options.Update(msg)
	return m, cmd
}

func (m MainModel) View() string {
	s := "Bubble Tea CLI Running...\n\n" + m.options.View()

	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s?", m.choice))
	}

	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Italic(true).
		Background(bgc).
		Foreground(fgc).
		Width(m.width).
		Height(m.height)

	return style.Render()
}
