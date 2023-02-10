package main

import (
	"fmt"
	"os"
	"strconv"
)

type Program struct {
	userIO               UserIO
	database             Database
	last_student         *Student
	transaction_executed int
}

func (program *Program) Program(userio UserIO, database Database) {
	program.userIO = userio
	program.database = database
	program.last_student = nil
	program.transaction_executed = 0
}

func (program *Program) run() {
	var menu uint

	program_menu := make(map[uint]string)
	program_menu[1] = "register student"
	program_menu[2] = "add course"
	program_menu[3] = "find student by id"
	program_menu[4] = "find student by name"
	program_menu[5] = "find students by course"
	program_menu[6] = "calculate gpa"
	program_menu[7] = "exit"

	var name string
	var id string
	var major string
	var age uint

	var course_name string
	var credit_hours uint
	var grade float32

	for {
		program.userIO.print("Enter Menu ID: ")
		fmt.Scan(&menu)

		if menu == 1 {

			program.userIO.print("Enter Student Name: ")
			name = program.userIO.input(name)

			program.userIO.print("Enter Student ID: ")
			id = program.userIO.input(id)

			program.userIO.print("Enter Student Major: ")
			major = program.userIO.input(major)

			program.userIO.print("Enter Student age: ")
			temp_age_string := program.userIO.input(strconv.FormatUint(uint64(age), 10))
			age_uint64, _ := strconv.ParseUint(temp_age_string, 10, 32)
			age = uint(age_uint64)

			student := program.database.add_student_with_all_params(name, id, major, age)

			program.last_student = student
			program.transaction_executed += 1
			program.userIO.print("Student " + name + " added\n")

		} else if menu == 2 {

			program.userIO.print("Enter Student ID: ")
			id = program.userIO.input(id)

			student, error := program.database.find_student_by_id(id)

			if error != nil {
				program.userIO.print("Student\n")
				program.userIO.print("\tName: " + student.name + "\n")
				program.userIO.print("\tID: " + student.id + "\n")
				program.userIO.print("\tMajor: " + student.major + "\n")
				program.userIO.print("\tAge: " + strconv.FormatUint(uint64(student.age), 10) + "\n")

				program.last_student = student
				program.transaction_executed += 1
			} else {
				program.userIO.print("Student not found\n")
			}
		} else if menu == 3 {
			program.userIO.print("Enter Student Name: ")
			name = program.userIO.input(name)

			student, error := program.database.find_student_by_name(name)

			if error != nil {
				program.userIO.print("Student\n")
				program.userIO.print("\tName: " + student.name + "\n")
				program.userIO.print("\tID: " + student.id + "\n")
				program.userIO.print("\tMajor: " + student.major + "\n")
				program.userIO.print("\tAge: " + strconv.FormatUint(uint64(student.age), 10) + "\n")

				program.last_student = student
				program.transaction_executed += 1
			} else {
				program.userIO.print("Student not found\n")
			}
		} else if menu == 4 {
			program.userIO.print("Enter Course Name: ")
			course_name = program.userIO.input(course_name)

			students := program.database.find_students_by_course(course_name)
			program.userIO.print("[")
			for _, student := range students {
				program.userIO.print(student.name + ",")
			}
			program.userIO.print("]\n")
			program.transaction_executed += 1

		} else if menu == 5 {

			if program.last_student != nil {
				program.userIO.print("Enter Course Name: ")
				course_name = program.userIO.input(course_name)

				program.userIO.print("Enter Course Credit Hours: ")
				temp_credit_hours_string := program.userIO.input(strconv.FormatUint(uint64(credit_hours), 10))
				credit_hours_uint64, _ := strconv.ParseUint(temp_credit_hours_string, 10, 32)
				credit_hours = uint(credit_hours_uint64)

				program.userIO.print("Enter Course Grade: ")
				temp_grade_string := program.userIO.input(strconv.FormatFloat(float64(grade), 'f', 2, 32))
				grade_float64, _ := strconv.ParseFloat(temp_grade_string, 32)
				grade = float32(grade_float64)

				program.last_student.add_course(course_name, credit_hours, grade)
				program.transaction_executed += 1
			} else {
				program.userIO.print("Search or add student before adding a course\n")
			}

		} else if menu == 6 {
			program.userIO.print("Transactions Executed: " + strconv.FormatUint(uint64(program.transaction_executed), 10) + "\n")
			program.transaction_executed += 1
		} else if menu == 7 {
			os.Exit(0)
		} else {
			program.userIO.print("Please enter a valid Menu option")
		}

	}

}
