package main

import (
	"github.com/luka2220/devtasks/cmd"
	"github.com/luka2220/devtasks/database"
)

func main() {
	database.OpenDBConnection()
	cmd.Execute()
}
