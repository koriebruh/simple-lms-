package repository

import "jamal/tgs/models/domain"

type CourseRepository interface {
	Create(course *domain.Course) (domain.Course, error)
	Delete(course *domain.Course, courseId int) error
	FindById(courseId int) (domain.Course, error)
	FindAll() ([]CourseFORMAT, error)
	Update(course *domain.Course, courseId int) (domain.Course, error)
}

// / INI UNTUK RERUTN DATA FINDALL (MENTAH BELUM ADA STATUS DAN HTTP STATUS)

type Teacher struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}

type CourseFORMAT struct {
	CourseID    uint    `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice uint    `json:"course_price"`
	Teacher     Teacher `json:"teacher"`
}
