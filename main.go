package main

import (
	"github.com/luka2220/devtasks/cmd"
	"github.com/luka2220/devtasks/database"
	// "github.com/luka2220/devtasks/database/seed"
)

func main() {
	database.OpenDBConnection()
	//seed.SeedBoards()
	cmd.Execute()
}
