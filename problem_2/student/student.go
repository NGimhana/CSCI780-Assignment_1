package student

type Course struct {
	CourseName  string
	CreditHours uint
	Grade       float32
}

func (student *Student) Add_course(course_name string, credit_hours uint, grade float32) *Student {
	course := Course{
		CourseName:  course_name,
		CreditHours: credit_hours,
		Grade:       grade,
	}
	student.Courses = append(student.Courses, course)
	return student
}

type Student struct {
	Name    string
	Id      string
	Major   string
	Age     uint
	Courses []Course
}

func (student *Student) NewStudent(name string, id string) *Student {
	student.Name = name
	student.Id = id
	student.Major = ""
	student.Age = 0
	student.Courses = []Course{}
	return student
}

func (student *Student) NewStudent_with_all_params(name string, id string, major string, age uint) *Student {
	student.Name = name
	student.Id = id
	student.Major = major
	student.Age = age
	student.Courses = []Course{}
	return student
}

func (student Student) Calculate_gpa() float32 {
	var gpa float32 = 0.0
	var sum float32 = 0.0
	var sum_credit_hours float32 = 0
	if len(student.Courses) > 0 {
		for _, course := range student.Courses {
			sum += course.Grade * float32(course.CreditHours)
			sum_credit_hours += float32(course.CreditHours)
		}
		gpa = sum / sum_credit_hours
	}
	return gpa
}
