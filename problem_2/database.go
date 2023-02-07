package main

import (
	"fmt"
)

type Database struct {
	students []Student
}

func (database *Database) Database() {
}

func (database *Database) add_student(name string, id string) *Student {
	student := Student{}
	student.NewStudent(name, id)
	database.students = append(database.students, student)
	fmt.Println(database.students[len(database.students)-1])
	return &database.students[len(database.students)-1]
}

func (database *Database) find_student_by_id(id string) Student {
	var searched_student *Student = nil
	for _, student := range database.students {
		if student.id == id {
			searched_student = &student
			break
		}
	}
	return *searched_student
}

func (database *Database) find_student_by_name(name string) Student {
	var searched_student *Student = nil
	for _, student := range database.students {
		if student.name == name {
			searched_student = &student
			break
		}
	}
	return *searched_student
}

func (database *Database) find_students_by_course(course_name string) []Student {
	var searched_students []Student = nil
	for _, student := range database.students {
		for _, course := range student.courses {
			if course.courseName == course_name {
				searched_students = append(searched_students, student)
				fmt.Println(searched_students)
				break
			}
		}
	}
	return searched_students
}

func (database *Database) num_student() uint {
	return uint(len(database.students))
}
