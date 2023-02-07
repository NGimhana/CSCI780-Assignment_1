package main

import "fmt"

type UserIO struct {
	message string
}

func (userio *UserIO) print(message string) {
	fmt.Println(message)
}

func (userio *UserIO) input(message string) {
	fmt.Println(message)
}

type Program struct {
	UserIO               UserIO
	database             Database
	last_student         Student
	transaction_executed int
}

func (program *Program) Program(userio UserIO, database Database) {
	program.UserIO = userio
	program.database = database
	program.last_student = Student{}
	program.transaction_executed = 0
}

func run() {
}
