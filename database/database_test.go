package database

import (
	"testing"
)

func TestCreateNewBoardDB(t *testing.T) {
	OpenTestDBConnection()

	res := CreateNewBoardDB("json-parser", true)

	if !res {
		t.Fail()
		t.Error("Board not stored in DB")
	}
}
