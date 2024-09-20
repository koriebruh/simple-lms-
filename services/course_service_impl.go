package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jamal/tgs/models/web"
	"jamal/tgs/repository"
	"net/http"
)

type CourseServiceImpl struct {
	courseRepository repository.CourseRepository
	db               *gorm.DB
	// latter we add validate
}

func (service CourseServiceImpl) Create(ctx *gin.Context) web.CourseCreateRequest {
	//TODO implement me

	/// PUNYA AKUN USER
	/// ROLES TEACHER OR ADMIN
	/// ADD DATA

	panic("implement me")
}

func (service CourseServiceImpl) Delete(ctx *gin.Context, courseId int) {
	//TODO implement me

	/// PUNYA AKUN USER
	/// ROLES TEACHER OR ADMIN
	/// FIND BY ID
	/// DELETE

	panic("implement me")
}

func (service CourseServiceImpl) FindById(ctx *gin.Context, courseId int) web.WebResponse {
	tx := service.db.Begin()
	if tx.Error != nil {
		return panic(errors.New("ERROR DI EKSEKUSI QUERY REPOSITORY"))
	}

	course, err := service.courseRepository.FindById(courseId)
	if err != nil {
		tx.Rollback()
	}

	tx.Commit()
	return web.WebResponse{
		Code: http.StatusOk,
	}

}

func (service CourseServiceImpl) FindAll(ctx *gin.Context) []web.WebResponse {
	//TODO implement me

	// FIND ALL

	panic("implement me")
}

func (service CourseServiceImpl) Update(ctx *gin.Context, courseId int) web.WebResponse {
	//TODO implement me

	/// PUNYA AKUN USER
	/// ROLES TEACHER OR ADMIN
	/// FIND BY ID
	/// UPDATE

	panic("implement me")
}
