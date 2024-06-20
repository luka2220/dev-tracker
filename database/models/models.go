package models

import "time"

// NOTE: Defining the fields for the Board database model
type Board struct {
	ID        uint      // Primary key & board id
	Name      string    // Hold the name for each board
	Active    bool      // Indicates if the board is active or not
	CreatedAt time.Time // Holds the creation time of each new model
	UpdatedAt time.Time // Holds the updated time of each new model
}

// NOTE: Defining fields for each Boards Task model
type Task struct {
	ID    uint   // Primary key & task id
	Tag   string // Holds the tag for each task
	Title string // Hold the title for each task
}
