package initialization

import (
	"fmt"
	"regexp"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/luka2220/devtasks/database"
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
	confirmationTextStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#FF8F00")).
				Italic(true)
	optionsTextStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#AF47D2"))
)

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
	state            int
	quitting         bool
	projectName      string
	setActiveBoard   bool
	errType          string
	errMessage       string
}

// NOTE: Initializes a new projectModel struct
// - Initializes new textinputs
// - Sets the default values for the projectModel struct
func StartInitModel() *projectModel {
	tiProject := textinput.New()
	tiProject.Prompt = ": "
	tiProject.Focus()
	tiProject.CharLimit = 30

	tiActive := textinput.New()
	tiActive.Prompt = ": "
	tiActive.CharLimit = 1

	return &projectModel{
		state:            0,
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
			// Reset the previous error type and message
			m.errType = ""
			m.errMessage = ""

			if m.state == 0 {
				projectNameValid := ValidateProjectNameInput(m.projectTi.Value())

				if projectNameValid {
					m.checkIfNameExists(m.projectTi.Value())
				} else {
					m.errType = "PROJECT_NAME"
					m.projectTi.Reset()
				}

			} else if m.state == 1 {
				switch m.setActiveBoardTi.Value() {
				case "y", "Y":
					m.handleActiveBoardInput(true)
				case "n", "N":
					m.handleActiveBoardInput(false)
				default:
					m.errType = "ACTIVE_BOARD"
					m.setActiveBoardTi.Reset()
				}
			} else if m.state == 2 {
				m.quitting = true
				return m, tea.Quit
			}
		}
	}

	m.projectTi, cmd = m.projectTi.Update(msg)
	m.setActiveBoardTi, cmd = m.setActiveBoardTi.Update(msg)

	return m, cmd
}

func (m *projectModel) View() string {
	s := fmt.Sprintf("Starting a New %s 💪.\n", infoTextStyle.Render("Development Kanban"))

	if m.state >= 0 {
		s += m.displayBoardNamePrompt()
	}

	if m.state >= 1 {
		s += m.displayActiveBoardPrompt()
	}

	if m.state >= 2 {
		s += m.displayBoardDatabaseStatus()
	}

	if m.quitting {
		s += "\nSee you next time! 👋\n"
	}

	return s
}

// Check whether the currently submitted board name already exists in the database
func (m *projectModel) checkIfNameExists(name string) {
	r, err := database.IsNameInDatabase(name)
	if err != nil {
		m.errType = "DATABASE"
		m.errMessage = err.Error()
	}

	if r {
		m.errType = "PROJECT_NAME_EXISTS"
		m.errMessage = name
		m.projectTi.Reset()
	} else {
		m.projectName = name
		m.state++
		m.projectTi.Blur()
		m.setActiveBoardTi.Focus()
	}
}

// Process the active board input from the prompt and store the created board
// inside the database
func (m *projectModel) handleActiveBoardInput(board bool) {
	m.setActiveBoard = board
	m.setActiveBoardTi.Blur()
	m.state += 1

	err := database.CreateNewBoardDB(m.projectName, m.setActiveBoard)
	if err != nil {
		m.errType = "DATABASE"
		m.errMessage = err.Error()
	}
}

// Displays and renders all of the prompts and logic for setting the new boards name
func (m *projectModel) displayBoardNamePrompt() string {
	var s string

	s = fmt.Sprintf("⚠️  %s use alpha-numeric characters or %s for new board names.\n",
		warningTextiStyle.Bold(true).Render("Only"),
		validTextStyle.Render("'-', '_', ' '"))

	if m.errType == "PROJECT_NAME" {
		s += fmt.Sprintf("❌ %s, only use alpha-numeric characters or %s for new board names.\n",
			errorTextStyle.Render("Invalid Name"),
			validTextStyle.Render("'-', '_', ' '"))
	} else if m.errType == "PROJECT_NAME_EXISTS" {
		s += fmt.Sprintf("❌ %s board name already exists in the database... Try another name.\n",
			errorTextStyle.Render(m.errMessage))
	} else if m.errType == "DATABASE" {
		s += fmt.Sprintf("❌ A %s occured. %s\n",
			errorTextStyle.Render("database error"), errorTextStyle.Render(m.errMessage))
	}

	s += fmt.Sprintf("What should the new board be called?%s\n", m.projectTi.View())

	return s
}

// Displays and renders all of the prompt logic for setting a board active or not
func (m *projectModel) displayActiveBoardPrompt() string {
	s := fmt.Sprintf("Created new %s to track development tasks!\n",
		validTextStyle.Render(m.projectName))

	if m.errType == "ACTIVE_BOARD" {
		return fmt.Sprintf("❌ %s, only use characters %s to indicate option.\n",
			errorTextStyle.Render("Invalid Input"),
			validTextStyle.Render("'y', 'Y', 'n', 'N'"))
	}

	s += fmt.Sprintf("Do you want to set the board %s as active❓ %s%s\n",
		confirmationTextStyle.Render(m.projectName),
		optionsTextStyle.Render("[y/n]"),
		m.setActiveBoardTi.View())

	if m.setActiveBoard {
		s += fmt.Sprintf("Set %s to active board!\n", validTextStyle.Render(m.projectName))
	}

	return s
}

// Displays successful storage of new board, or any errors
func (m *projectModel) displayBoardDatabaseStatus() string {
	var s string
	if m.errType == "DATABASE" {
		return fmt.Sprintf("❌ A %s occured. %s\n",
			errorTextStyle.Render("database error"), errorTextStyle.Render(m.errMessage))
	}
	s = fmt.Sprintf("%s record created in db. The board is good to use!\n", validTextStyle.Render(m.projectName))

	return s
}

// Validate input for project name with regular expressions
func ValidateProjectNameInput(input string) bool {
	pattern := "^[a-zA-Z0-9_\\- ]+$"
	re := regexp.MustCompile(pattern)

	return re.MatchString(input)
}
