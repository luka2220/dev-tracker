package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/luka2220/devtasks/constants"
)

var (
	infoTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#3572EF"))
	warningTextiStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FFBF00"))
	validTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#059212"))
	confirmationTextiStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF8F00")).
				Italic(true)
	optionsTextStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#AF47D2"))
)

func StartProjectInitTui() {
	m := initializeModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		constants.Logger.WriteString(fmt.Sprintf("Error starting project init tui: %v", err))
		os.Exit(1)
	}
}

type projectModelDB struct {
	name   string
	active bool
}

type projectModel struct {
	count            int // holds the input state
	quitting         bool
	projectTextInput textinput.Model
	projectName      string
	activeTextInput  textinput.Model
	active           bool
	err              error
}

func initializeModel() *projectModel {
	tiProject := textinput.New()
	tiProject.Prompt = ": "
	tiProject.Focus()
	tiProject.CharLimit = 30

	tiActive := textinput.New()
	tiActive.Prompt = ": "
	tiActive.CharLimit = 1

	return &projectModel{
		count:            0,
		quitting:         false,
		projectTextInput: tiProject,
		activeTextInput:  tiActive,
		err:              nil,
	}
}

func (m *projectModel) Init() tea.Cmd {
	return nil
}

func (m *projectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "enter":
			if m.count == 0 {
				m.projectName = m.projectTextInput.Value()
				m.count++
				m.projectTextInput.Blur()
				m.activeTextInput.Focus()
			} else if m.count == 1 {
				switch m.activeTextInput.Value() {
				case "y":
					m.active = true
				case "n":
					m.active = false
				}
				m.activeTextInput.Blur()
				m.count++

				return m, tea.Quit
			}
		}
	}

	m.projectTextInput, cmd = m.projectTextInput.Update(msg)
	m.activeTextInput, cmd = m.activeTextInput.Update(msg)

	return m, cmd
}

func (m *projectModel) View() string {
	s := fmt.Sprintf("Starting a New %s 💪.\n", infoTextStyle.Render("Development Kanban"))
	s += fmt.Sprintf("⚠️  %s use alpha-numeric characters or %s for new board names.\n",
		warningTextiStyle.Bold(true).Render("Only"),
		validTextStyle.Render("'-', '_', ' '"))

	s += fmt.Sprintf("What should the new board be called?%s\n", m.projectTextInput.View())

	if m.count >= 1 {
		s += fmt.Sprintf("Created new %s to track development tasks!\n",
			validTextStyle.Render(m.projectName))

		s += fmt.Sprintf("Do you want to set the board %s as active❓ %s%s\n",
			confirmationTextiStyle.Render(m.projectName),
			optionsTextStyle.Render("[y/n]"),
			m.activeTextInput.View())

	}

	if m.quitting {
		s += "\nSee you next time! 👋\n"
	}

	return s
}
