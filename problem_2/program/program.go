package program

import (
	"fmt"
	"os"
	"problem_2/database"
	"problem_2/student"
	"problem_2/userIO"
	"strconv"
)

type Program struct {
	userIO               userIO.UserIO
	database             database.Database
	last_student         *student.Student
	transaction_executed int
}

func (program *Program) Program(userio userIO.UserIO, database database.Database) {
	program.userIO = userio
	program.database = database
	program.last_student = nil
	program.transaction_executed = 0
}

func (program *Program) Run() {
	var menu uint

	program_menu := make(map[uint]string)
	program_menu[1] = "Add a student"
	program_menu[2] = "Find a student by id"
	program_menu[3] = "Find a student by name"
	program_menu[4] = "List students by course"
	program_menu[5] = "Add a course to a student"
	program_menu[6] = "Print the number of transactions executed"
	program_menu[7] = "exit"

	var name string
	var id string
	var major string
	var age uint

	var course_name string
	var credit_hours uint
	var grade float32

	for {

		program.userIO.Print("\nSelect a Menu from the below list: \n")
		for key, value := range program_menu {
			program.userIO.Print(strconv.FormatUint(uint64(key), 10) + ". " + value + "\n")
		}

		program.userIO.Print("Enter Menu ID: ")
		fmt.Scan(&menu)

		if menu == 1 {

			program.userIO.Print("Enter Student Name: ")
			name = program.userIO.Input(name)

			program.userIO.Print("Enter Student ID: ")
			id = program.userIO.Input(id)

			program.userIO.Print("Enter Student Major: ")
			major = program.userIO.Input(major)

			program.userIO.Print("Enter Student age: ")
			temp_age_string := program.userIO.Input(strconv.FormatUint(uint64(age), 10))
			age_uint64, _ := strconv.ParseUint(temp_age_string, 10, 32)
			age = uint(age_uint64)

			student := program.database.Add_student_with_all_params(name, id, major, age)

			program.last_student = student
			program.transaction_executed += 1
			program.userIO.Print("Student " + name + " added\n")

		} else if menu == 2 {

			program.userIO.Print("Enter Student ID: ")
			id = program.userIO.Input(id)

			student := program.database.Find_student_by_id(id)

			if student != nil {
				program.userIO.Print("Student\n")
				program.userIO.Print("\tName: " + student.Name + "\n")
				program.userIO.Print("\tID: " + student.Id + "\n")
				program.userIO.Print("\tMajor: " + student.Major + "\n")
				program.userIO.Print("\tAge: " + strconv.FormatUint(uint64(student.Age), 10) + "\n")

				program.last_student = student
				program.transaction_executed += 1
			} else {
				program.userIO.Print("Student not found\n")
			}
		} else if menu == 3 {
			program.userIO.Print("Enter Student Name: ")
			name = program.userIO.Input(name)

			student := program.database.Find_student_by_name(name)

			if student != nil {
				program.userIO.Print("Student\n")
				program.userIO.Print("\tName: " + student.Name + "\n")
				program.userIO.Print("\tID: " + student.Id + "\n")
				program.userIO.Print("\tMajor: " + student.Major + "\n")
				program.userIO.Print("\tAge: " + strconv.FormatUint(uint64(student.Age), 10) + "\n")

				program.last_student = student
				program.transaction_executed += 1
			} else {
				program.userIO.Print("Student not found\n")
			}
		} else if menu == 4 {
			program.userIO.Print("Enter Course Name: ")
			course_name = program.userIO.Input(course_name)

			students := program.database.Find_students_by_course(course_name)
			if len(students) != 0 {
				program.userIO.Print("[")
				for _, student := range students {
					program.userIO.Print(student.Name + ",")
				}
				program.userIO.Print("\b")
				program.userIO.Print("]\n")
				program.transaction_executed += 1
			} else {
				program.userIO.Print("No Student found\n")
			}
		} else if menu == 5 {

			if program.last_student != nil {
				program.userIO.Print("Enter Course Name: ")
				course_name = program.userIO.Input(course_name)

				program.userIO.Print("Enter Course Credit Hours: ")
				temp_credit_hours_string := program.userIO.Input(strconv.FormatUint(uint64(credit_hours), 10))
				credit_hours_uint64, _ := strconv.ParseUint(temp_credit_hours_string, 10, 32)
				credit_hours = uint(credit_hours_uint64)

				program.userIO.Print("Enter Course Grade: ")
				temp_grade_string := program.userIO.Input(strconv.FormatFloat(float64(grade), 'f', 2, 32))
				grade_float64, _ := strconv.ParseFloat(temp_grade_string, 32)
				grade = float32(grade_float64)

				program.last_student.Add_course(course_name, credit_hours, grade)
				program.transaction_executed += 1

				program.userIO.Print("Course " + course_name + " added to student " + program.last_student.Name + "\n")

			} else {
				program.userIO.Print("Search or add student before adding a course\n")
			}

		} else if menu == 6 {
			program.userIO.Print("Transactions Executed: " + strconv.FormatUint(uint64(program.transaction_executed), 10) + "\n")
			program.transaction_executed += 1
		} else if menu == 7 {
			os.Exit(0)
		} else {
			program.userIO.Print("Please enter a valid Menu option")
		}

	}

}
