package main

import (
	"sort"
)

type Database struct {
	students         []Student
	index_to_student map[int]Student
}

func (database *Database) Database() {
	database.index_to_student = make(map[int]Student)
}

func (database *Database) add_student(name string, id string) *Student {
	student := Student{}
	student.NewStudent(name, id)
	index := database.num_student()
	database.index_to_student[int(index)+1] = student
	database.students = append(database.students, student)
	return &database.students[len(database.students)-1]
}

func (database *Database) add_student_with_all_params(name string, id string, major string, age uint) *Student {
	student := Student{}
	student.NewStudent_with_all_params(name, id, major, age)
	index := database.num_student()
	database.index_to_student[int(index)+1] = student
	database.students = append(database.students, student)
	return &database.students[len(database.students)-1]
}

func (database *Database) find_student_by_id(id string) (*Student, error) {
	searched_student_index := 0
	for index, student := range database.students {
		if student.id == id {
			searched_student_index = index
			return &database.students[searched_student_index], nil
		}
	}
	return nil, nil

}

func (database *Database) find_student_by_name(name string) (*Student, error) {
	searched_student_index := 0
	for index, student := range database.students {
		if student.name == name {
			searched_student_index = index
			return &database.students[searched_student_index], nil
		}
	}
	return nil, nil

}

func (database *Database) find_students_by_course(course_name string) []Student {
	var searched_students []Student = nil
	keys := make([]int, 0, len(database.index_to_student))

	for k := range database.index_to_student {
		keys = append(keys, k)
	}
	// this will give the student items in the same order they were added
	sort.Ints(keys)

	// iterate by sorted keys
	for _, index := range keys {
		student, _ := database.find_student_by_id(database.index_to_student[index].id)
		for _, course := range student.courses {
			if course.courseName == course_name {
				searched_students = append(searched_students, *student)
				break
			}
		}
	}

	return searched_students
}

func (database *Database) num_student() uint {
	return uint(len(database.students))
}
