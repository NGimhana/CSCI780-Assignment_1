package database

import (
	"problem_2/student"
	"sort"
)

type Database struct {
	students         []student.Student
	index_to_student map[int]student.Student
}

func (database *Database) Database() {
	database.index_to_student = make(map[int]student.Student)
}

func (database *Database) Add_student(name string, id string) *student.Student {
	student := student.Student{}
	student.NewStudent(name, id)
	index := database.num_student()
	database.index_to_student[int(index)+1] = student
	database.students = append(database.students, student)
	return &database.students[len(database.students)-1]
}

func (database *Database) Add_student_with_all_params(name string, id string, major string, age uint) *student.Student {
	student := student.Student{}
	student.NewStudent_with_all_params(name, id, major, age)
	index := database.num_student()
	database.index_to_student[int(index)+1] = student
	database.students = append(database.students, student)
	return &database.students[len(database.students)-1]
}

func (database *Database) Find_student_by_id(id string) *student.Student {
	searched_student_index := 0
	for index, student := range database.students {
		if student.Id == id {
			searched_student_index = index
			return &database.students[searched_student_index]
		}
	}
	return nil

}

func (database *Database) Find_student_by_name(name string) *student.Student {
	searched_student_index := 0
	for index, student := range database.students {
		if student.Name == name {
			searched_student_index = index
			return &database.students[searched_student_index]
		}
	}
	return nil

}

func (database *Database) Find_students_by_course(course_name string) []student.Student {
	var searched_students []student.Student = nil
	keys := make([]int, 0, len(database.index_to_student))

	for k := range database.index_to_student {
		keys = append(keys, k)
	}
	// this will give the student items in the same order they were added
	sort.Ints(keys)

	// iterate by sorted keys
	for _, index := range keys {
		student := database.Find_student_by_id(database.index_to_student[index].Id)
		for _, course := range student.Courses {
			if course.CourseName == course_name {
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
