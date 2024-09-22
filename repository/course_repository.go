package repository

import "jamal/tgs/models/domain"

type CourseRepository interface {
	Create(course *domain.Course) (domain.Course, error)
	Delete(course *domain.Course, courseId int) error
	FindById(courseId int) (domain.Course, error)
	FindAll() ([]domain.CourseFindAll, error)
	Update(course *domain.Course, courseId int) (domain.Course, error)
}

// / INI UNTUK RERUTN DATA FINDALL (MENTAH BELUM ADA STATUS DAN HTTP STATUS)
