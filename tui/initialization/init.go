package initialization

import (
	"fmt"
	"os"
	"regexp"

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
	errorTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#C80036"))
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
// setActiveBoardTi (textinput.Model): textinput element for setting board active
// count (int): state for tracking which textinput the user is currently on
// quitting (bool): state for quitting out of the CLI
// projectName (string): stores the result of projectTi from user
// setActiveBoard (bool): stores the result of setActiveBoardTi from user
type projectModel struct {
	projectTi        textinput.Model
	setActiveBoardTi textinput.Model
	count            int
	quitting         bool
	projectName      string
	setActiveBoard   bool
	errMessage       string
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

	return &projectModel{
		count:            0,
		quitting:         false,
		projectTi:        tiProject,
		setActiveBoardTi: tiActive,
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
				projectNameValid := ValidateProjectNameInput(m.projectTi.Value())

				if projectNameValid {
					m.projectName = m.projectTi.Value()
					m.count++
					m.projectTi.Blur()
					m.setActiveBoardTi.Focus()
				} else {
					m.errMessage = "PROJECT_NAME"
					m.projectTi.Reset()
				}

			} else if m.count == 1 {
				switch m.setActiveBoardTi.Value() {
				case "y", "Y":
					m.setActiveBoard = true
					m.setActiveBoardTi.Blur()
					return m, tea.Quit

				case "n", "N":
					m.setActiveBoard = false
					m.setActiveBoardTi.Blur()
					return m, tea.Quit

				default:
					m.errMessage = "ACTIVE_BOARD"
					m.setActiveBoardTi.Reset()
				}
			}
		}
	}

	m.projectTi, cmd = m.projectTi.Update(msg)
	m.setActiveBoardTi, cmd = m.setActiveBoardTi.Update(msg)

	return m, cmd
}

func (m *projectModel) View() string {
	s := fmt.Sprintf("Starting a New %s 💪.\n", infoTextStyle.Render("Development Kanban"))
	s += fmt.Sprintf("⚠️  %s use alpha-numeric characters or %s for new board names.\n",
		warningTextiStyle.Bold(true).Render("Only"),
		validTextStyle.Render("'-', '_', ' '"))

	if m.errMessage == "PROJECT_NAME" {
		s += fmt.Sprintf("❌ %s, only use alpha-numeric characters or %s for new board names.\n",
			errorTextStyle.Render("Invalid Name"),
			validTextStyle.Render("'-', '_', ' '"))
	}

	s += fmt.Sprintf("What should the new board be called?%s\n", m.projectTi.View())

	if m.count == 1 {
		s += fmt.Sprintf("Created new %s to track development tasks!\n",
			validTextStyle.Render(m.projectName))

		if m.errMessage == "ACTIVE_BOARD" {
			s += fmt.Sprintf("❌ %s, only use characters %s to indicate option.\n",
				errorTextStyle.Render("Invalid Input"),
				validTextStyle.Render("'y', 'Y', 'n', 'N'"))
		}

		s += fmt.Sprintf("Do you want to set the board %s as active❓ %s%s\n",
			confirmationTextiStyle.Render(m.projectName),
			optionsTextStyle.Render("[y/n]"),
			m.setActiveBoardTi.View())
	}

	if m.setActiveBoard {
		s += fmt.Sprintf("Set %s to active board!\n", validTextStyle.Render(m.projectName))
	}

	if m.quitting {
		s += "\nSee you next time! 👋\n"
	}

	return s
}

func ValidateProjectNameInput(input string) bool {
	pattern := "^[a-zA-Z0-9_\\- ]+$"
	re := regexp.MustCompile(pattern)

	return re.MatchString(input)
}
