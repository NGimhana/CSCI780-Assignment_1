package main

import "fmt"

func main() {
	database := Database{}
	database.Database()

	stu2 := database.add_student("John doe", "1234")
	stu2.add_course("Math", 2, 3.5)

	stu1 := database.add_student("John", "123")
	stu1.add_course("Math", 3, 3.5)

	stu3 := database.add_student("Anna", "12345")
	stu3.add_course("english", 2, 3.5)

	fmt.Println(database.find_students_by_course("Math"))

}
