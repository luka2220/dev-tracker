package database

import (
	"testing"
)

func TestCreateNewBoardDB(t *testing.T) {
	OpenTestDBConnection()

	// TODO: Write test cases based on results from DB
	result := CreateNewBoardDB("dev-tui", true)
	if !result {
		t.Fail()
		t.Error("Board not stored in DB")
	}
}
