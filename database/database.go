package database

import (
	"fmt"
	"time"

	//"github.com/luka2220/devtasks/utils"
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

	fmt.Println("Successfully opened DB")

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
// TODO: Add better error handling for caller
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

	return nil
}

func CreateNewBoardDBConcurrent() {}
