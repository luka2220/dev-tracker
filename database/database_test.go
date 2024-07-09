package database

import (
	"testing"
)

func TestCreateNewBoardDB(t *testing.T) {
	OpenTestDBConnection()

	// Tests the function for creating a new developmentboard
	t1 := CreateNewBoardDB("dev-tui", true)
	if t1 != nil {
		t.Fail()
		t.Error("Board not stored in DB")
	}

}

func TestCreateDevelopmentTask(t *testing.T) {
	OpenDBConnection()

	// Test the function for createing a new development task
}
