package main

func main() {

	userIO := UserIO{}
	database := Database{}
	database.Database()

	program := Program{}
	program.Program(userIO, database)

	program.run()
}
