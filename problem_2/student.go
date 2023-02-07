package main

type Course struct {
	courseName  string
	creditHours uint
	grade       float32
}

func (student *Student) add_course(course_name string, credit_hours uint, grade float32) {
	course := Course{
		courseName:  course_name,
		creditHours: credit_hours,
		grade:       grade,
	}
	student.courses = append(student.courses, course)
}

type Student struct {
	name    string
	id      string
	major   string
	age     int
	courses []Course
}

func (student *Student) NewStudent(name string, id string) *Student {
	student.name = name
	student.id = id
	student.major = ""
	student.age = 0
	student.courses = []Course{}
	return student
}

func (student Student) calculate_gpa() float32 {
	var gpa float32 = 0.0
	var sum float32 = 0.0
	var sum_credit_hours float32 = 0
	if len(student.courses) > 0 {
		for _, course := range student.courses {
			sum += course.grade * float32(course.creditHours)
			sum_credit_hours += float32(course.creditHours)
		}
		gpa = sum / sum_credit_hours
	}
	return gpa
}
