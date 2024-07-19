package database

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// NOTE: Defining the fields for the Board database model
// A Board can have many tasks (one to many relationship)
type Board struct {
	gorm.Model
	Tasks     []Task    `gorm:"foreignKey:TaskId"` // All tasks for the board
	ID        uint      `gorm:"primaryKey"`        // Primary key & board id
	Name      string    // Hold the name for each board
	Active    bool      // Indicates if the board is active or not
	CreatedAt time.Time // Holds the creation time of each new model
	UpdatedAt time.Time // Holds the updated time of each new model
}

// NOTE: Defining fields for each Boards Task model
// A Task can only have one board (one to many relationship)
type Task struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"` // Primary key & task id
	Tag         string // Holds the tag for each task
	Title       string // Hold the title for each task
	Description string // Full description and text of task
	State       string // Current state of the task. Can be Ready|Progress|Done
	DaysLeft    int    // Amount of days left until due
	TaskId      uint   // Foreign key id
}

// NOTE: Opens connection to the sqlite db
// Throws and error if connection is failed
func OpenDBConnection() {
	dbconn, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Error opening connection to the database")
	}

	dbconn.AutoMigrate(&Board{}, &Task{})
	db = dbconn
}

// NOTE: Opens a connection for testing the database
func OpenTestDBConnection() {
	dbconn, err := gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
	if err != nil {
		panic("Error opening connection to the database")
	}

	dbconn.AutoMigrate(&Board{}, &Task{})

	fmt.Println("Successfully opened DB")

	db = dbconn
}

// NOTE: Creates a new board record in the db
func CreateNewBoardDB(name string, active bool) error {
	var boards []Board

	records := db.Raw("SELECT id, name, active FROM boards").Find(&boards)
	if records.Error != nil {
		return fmt.Errorf("An error occured getting records from the db: %v", records.Error)
	}

	for _, board := range boards {
		if board.Name == name {
			return fmt.Errorf("Board name already exists in db...")
		}

		if active && board.Active {
			db.Save(&Board{ID: board.ID, Name: board.Name, Active: false})
		}
	}

	result := db.Create(&Board{Name: name, Active: active})
	if result.Error != nil {
		return fmt.Errorf("An error occured creating a new board record in the db: %v", result.Error)
	}

	// Log the newly created board name and active state
	LogDBSlog(fmt.Sprintf("New board created. name=%s, active=%v", name, active))

	return nil
}

// NOTE: Checks if the given board name exists in the database
func IsNameInDatabase(name string) (bool, error) {
	var boards []Board

	records := db.Raw("SELECT name FROM boards").Find(&boards)
	if records.Error != nil {
		return false, fmt.Errorf("Error getting board names from database: %v", records.Error)
	}

	for _, board := range boards {
		if board.Name == name {
			return true, nil
		}
	}

	return false, nil
}

// NOTE: Returns the currently active board from the db
func GetActiveBoard() (Board, error) {
	var board Board

	record := db.First(&board, "active = ?", "1")
	if record.Error != nil {
		return board, fmt.Errorf("Error getting active record from database: %v", record.Error)
	}

	return board, nil
}

// NOTE: Creates a new task in the currently active board
func CreateNewDevelopmentTask() {}

// NOTE: Database logging helper
func LogDBSlog(msg string) {
	file, err := os.OpenFile("./database/db.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		msg := fmt.Sprintf("Error opening db.log file: %v", err)
		panic(msg)
	}

	defer file.Close()

	logger := slog.New(slog.NewTextHandler(file, &slog.HandlerOptions{}))
	slog.SetDefault(logger)

	logMsg := msg + " (slog)"
	slog.Info(logMsg)
}
