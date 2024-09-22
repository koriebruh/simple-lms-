package repository

import (
	"fmt"
	"gorm.io/gorm"
	"jamal/tgs/models/domain"
)

type CourseRepositoryImpl struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *CourseRepositoryImpl {
	return &CourseRepositoryImpl{db: db}
}

func (repository CourseRepositoryImpl) Create(course *domain.Course) (domain.Course, error) {
	result := repository.db.Create(course)
	if result.Error != nil {
		return domain.Course{}, result.Error
	}

	return *course, nil
}

func (repository CourseRepositoryImpl) Delete(course *domain.Course, courseId int) error {
	/// === DELETE FROM courses WHERE id = courseId;
	result := repository.db.Delete(course, courseId)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository CourseRepositoryImpl) FindById(courseId int) (domain.Course, error) {
	var course domain.Course
	/// SELECT * FROM courses WHERE id = 1 LIMIT 1;
	result := repository.db.Where("id = ?", courseId).First(&course)

	if result.Error != nil {
		return domain.Course{}, result.Error
	}

	return course, nil
}

func (repository CourseRepositoryImpl) FindAll() ([]domain.CourseFindAll, error) {
	var results []struct {
		CourseID    uint   `json:"course_id"`
		CourseName  string `json:"course_name"`
		CoursePrice uint   `json:"course_price"`
		UserID      uint   `json:"user_id"`
		Username    string `json:"username"`
		FullName    string `json:"full_name"`
	}

	QUERY := `
		SELECT courses.id AS course_id,
		       courses.name AS course_name,
		       courses.price AS course_price,
		       users.id AS user_id,
		       users.username,
		       users.full_name
		FROM courses
		INNER JOIN users ON courses.teacher = users.id
	`

	result := repository.db.Raw(QUERY).Scan(&results)
	if result.Error != nil {
		return nil, result.Error
	}

	courses := make([]domain.CourseFindAll, len(results))
	for i, r := range results {
		courses[i] = domain.CourseFindAll{
			CourseID:    r.CourseID,
			CourseName:  r.CourseName,
			CoursePrice: r.CoursePrice,
			Teacher: domain.Teacher{
				UserID:   r.UserID,
				Username: r.Username,
				FullName: r.FullName,
			},
		}
	}

	return courses, nil

}

func (repository CourseRepositoryImpl) Update(course *domain.Course, courseId int) (domain.Course, error) {
	var CourseNil domain.Course

	// <-- UPDATE data dan data di masukan ke variable course
	result := repository.db.Model(&domain.Course{}).Where("id = ?", courseId).Updates(course)

	if result.Error != nil {
		return CourseNil, result.Error
	}

	if result.RowsAffected == 0 { //<-- jika tidak ada baris yang terupdate maka akan err
		return CourseNil, fmt.Errorf("course with id %d not found", courseId)
	}

	return *course, nil
}
