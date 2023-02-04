package student

type Course struct {
	courseName  string
	creditHours uint
	grade       float32
}

type Student struct {
	name    string
	id      string
	major   string
	age     int
	courses []Course
}

func NewStudent(name string, id string) Student {

	student := Student{
		name:    name,
		id:      id,
		major:   "",
		age:     0,
		courses: []Course{},
	}
	return student
}

func add_course(student Student, course_name string, credit_hours uint, grade float32) Course {
	course := Course{
		courseName:  course_name,
		creditHours: credit_hours,
		grade:       grade,
	}
	return course

}
