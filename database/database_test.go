package database

import (
	"testing"
)

func TestCreateNewBoardDB(t *testing.T) {
	OpenTestDBConnection()

	// TODO: Write test cases based on results from DB
	t1 := CreateNewBoardDB("dev-tui", true)
	if t1 != nil {
		t.Fail()
		t.Error("Board not stored in DB")
	}
}
