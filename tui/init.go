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

// NOTE: Project Initialization Type
// projectTi (textinput.Model): textinput element for the project name
// setGhRepoOptionTi (textinput.Model): textinput element for option to add a github repo
// ghRepoTi (textinput.Model): textinput element for the github repo
// setActiveBoardTi (textinput.Model): textinput element for setting board active
// count (int): state for tracking which textinput the user is currently on
// quitting (bool): state for quitting out of the CLI
// projectName (string): stores the result of projectTi from user
// setActiveBoard (bool): stores the result of setActiveBoardTi from user
// setGithubRepo (bool): stores the result of setGhRepoOptionTi from user
// githubRepo (string): stores the result of ghRepoTi from user

type projectModel struct {
	projectTi         textinput.Model
	setGhRepoOptionTi textinput.Model
	ghRepoTi          textinput.Model
	setActiveBoardTi  textinput.Model
	count             int
	quitting          bool
	projectName       string
	setActiveBoard    bool
	setGithubRepo     bool
	githubRepo        string
	err               error
}

// NOTE: Initializes a new projectModel struct
// - Initializes new textinputs
// - Sets the default values for the projectModel struct

func initializeModel() *projectModel {
	tiProject := textinput.New()
	tiProject.Prompt = ": "
	tiProject.Focus()
	tiProject.CharLimit = 30

	tiActive := textinput.New()
	tiActive.Prompt = ": "
	tiActive.CharLimit = 1

	tiGHRepoOption := textinput.New()
	tiGHRepoOption.Prompt = ": "
	tiGHRepoOption.CharLimit = 1

	tiGHRepo := textinput.New()
	tiGHRepo.Prompt = ": "
	tiGHRepo.CharLimit = 50

	return &projectModel{
		count:             0,
		quitting:          false,
		projectTi:         tiProject,
		setActiveBoardTi:  tiActive,
		setGhRepoOptionTi: tiGHRepoOption,
		ghRepoTi:          tiGHRepo,
		err:               nil,
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
				m.projectName = m.projectTi.Value()
				m.count++
				m.projectTi.Blur()
				m.setActiveBoardTi.Focus()
			} else if m.count == 1 {
				switch m.setActiveBoardTi.Value() {
				case "y":
					m.setActiveBoard = true
				case "n":
					m.setActiveBoard = false
				}
				m.setActiveBoardTi.Blur()
				m.count++

				return m, tea.Quit
			}
		}
	}

	m.projectTi, cmd = m.projectTi.Update(msg)
	m.setActiveBoardTi, cmd = m.setActiveBoardTi.Update(msg)
	m.setGhRepoOptionTi, cmd = m.setGhRepoOptionTi.Update(msg)
	m.ghRepoTi, cmd = m.ghRepoTi.Update(msg)

	return m, cmd
}

func (m *projectModel) View() string {
	s := fmt.Sprintf("Starting a New %s 💪.\n", infoTextStyle.Render("Development Kanban"))
	s += fmt.Sprintf("⚠️  %s use alpha-numeric characters or %s for new board names.\n",
		warningTextiStyle.Bold(true).Render("Only"),
		validTextStyle.Render("'-', '_', ' '"))

	s += fmt.Sprintf("What should the new board be called?%s\n", m.projectTi.View())

	if m.count >= 1 {
		s += fmt.Sprintf("Created new %s to track development tasks!\n",
			validTextStyle.Render(m.projectName))

		s += fmt.Sprintf("Do you want to set the board %s as active❓ %s%s\n",
			confirmationTextiStyle.Render(m.projectName),
			optionsTextStyle.Render("[y/n]"),
			m.setActiveBoardTi.View())

	}

	if m.quitting {
		s += "\nSee you next time! 👋\n"
	}

	return s
}
