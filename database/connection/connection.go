package utils

import (
	"github.com/luka2220/devtasks/constants"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func OpenDBConnection() {
	dbconn, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		constants.Logger.WriteString("(connection.go 14) Error openning database")
		panic("Error opening connection to the database")
	}

	DBConn = dbconn
}
