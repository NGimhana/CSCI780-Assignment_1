package main

import "fmt"

func main() {

	student := Student{}
	student.NewStudent("John", "12345")
	student.add_course("Math", 3, 3.5)

	database := Database{}
	database.NewDatabase()

	stu1 := database.add_student("John", "12345")
	database.find_student_by_id("12345").add_course("Math", 3, 3.5)

	stu2 := database.add_student("John Doe", "12346")
	stu2.add_course("English", 1, 1.5)

	stu3 := database.add_student("Anna", "12347")
	stu3.add_course("Math", 2, 2.5)

	fmt.Println(database.find_students_by_course("Math1"))
}
