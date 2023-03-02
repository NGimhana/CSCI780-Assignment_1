package main

import (
	"problem_2/database"
	"problem_2/program"
	"problem_2/userIO"
)

func main() {

	userIO := userIO.UserIO{}
	database := database.Database{}
	database.Database()

	program := program.Program{}
	program.Program(userIO, database)

	program.Run()
}
