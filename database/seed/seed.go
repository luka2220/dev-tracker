package seed

import (
	"github.com/luka2220/devtasks/database"
)

func SeedBoards() {
	database.CreateNewBoardDB("interpreter", false)
	database.CreateNewBoardDB("json-parser", false)
	database.CreateNewBoardDB("compression-tool", false)
	database.CreateNewBoardDB("pdf-tool", false)
	database.CreateNewBoardDB("redis-cli", false)
}
