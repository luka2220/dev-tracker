package database

import (
	"fmt"
	"time"

	"github.com/luka2220/devtasks/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// NOTE: Defining the fields for the Board database model
// A board can have many tasks (one to many relationship)
type board struct {
	gorm.Model
	Tasks     []task    `gorm:"foreignKey:TaskId"` // All tasks for the board
	ID        uint      `gorm:"primaryKey"`        // Primary key & board id
	Name      string    // Hold the name for each board
	Active    bool      // Indicates if the board is active or not
	CreatedAt time.Time // Holds the creation time of each new model
	UpdatedAt time.Time // Holds the updated time of each new model
}

// NOTE: Defining fields for each Boards Task model
// A task can only have one board (one to many relationship)
type task struct {
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

	dbconn.AutoMigrate(&board{}, &task{})

	fmt.Println("Successfully opened DB")

	db = dbconn
}

// NOTE: Opens a connection for testing the database
func OpenTestDBConnection() {
	dbconn, err := gorm.Open(sqlite.Open("../database.db"), &gorm.Config{})
	if err != nil {
		panic("Error opening connection to the database")
	}

	dbconn.AutoMigrate(&board{}, &task{})

	fmt.Println("Successfully opened DB")

	db = dbconn
}

// NOTE: Creates a new board record in the db

// TODO:Function Operations
// - Check if the DB is empty (if so make the board active)
// - Figure out the id?
// - Store the record in the DB
// - Return a bool based on the store operation result
func CreateNewBoardDB(name string, active bool) bool {

	if active {
		var activeBoard board

		// Retrieve the currently active board
		db.Raw("SELECT * FROM boards WHERE active = ?", 1).Find(&activeBoard)

		var log string

		// Check if any result were returned from the db
		if activeBoard.ID == 0 {
			log = fmt.Sprintf("No records have an active board")
		} else {
			log = fmt.Sprintf("id=%d, name=%s, active=%v", activeBoard.ID, activeBoard.Name, activeBoard.Active)
		}

		utils.LogDB(log)
	}

	return true
}
