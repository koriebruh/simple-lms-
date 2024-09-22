package services

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"jamal/tgs/helper"
	"jamal/tgs/models/web"
	"jamal/tgs/repository"
	"net/http"
)

type CourseServiceImpl struct {
	courseRepository repository.CourseRepository
	db               *gorm.DB
}

func NewCourseServiceImpl(courseRepo repository.CourseRepository, db *gorm.DB) CourseService {
	return &CourseServiceImpl{
		courseRepository: courseRepo,
		db:               db,
	}
}

func (service CourseServiceImpl) Create(ctx *gin.Context) web.CourseCreateRequest {
	//TODO implement me

	/// PUNYA AKUN USER
	/// ROLES TEACHER OR ADMIN
	/// ADD DATA

	panic("implement me")
}

func (service CourseServiceImpl) Delete(ctx *gin.Context, courseId int) web.WebResponse {
	/// START TRANSACTION
	var err error
	tx := service.db.Begin()
	if tx.Error != nil {
		return helper.WebResponsesFunc(http.StatusInternalServerError, "INTERNAL SERVER ERROR", "error in start tx")
	}
	defer helper.CommitOrRollback(tx, &err)

	/// ROLES TEACHER OR ADMIN

	/// FIND BY ID
	/// DELETE
	panic("implement me")
}

func (service CourseServiceImpl) FindById(ctx *gin.Context, courseId int) web.WebResponse {
	/// START TRANSACTION
	var err error
	tx := service.db.Begin()
	if tx.Error != nil {
		return helper.WebResponsesFunc(http.StatusInternalServerError, "INTERNAL SERVER ERROR", "error in start tx")
	}
	defer helper.CommitOrRollback(tx, &err)

	/// FIND ID
	course, err := service.courseRepository.FindById(courseId)
	if err != nil {
		return helper.WebResponsesFunc(http.StatusBadRequest, "BAD REQUEST", "id not found")
	}

	return helper.WebResponsesFunc(http.StatusOK, "OK", course)

}

func (service CourseServiceImpl) FindAll(ctx *gin.Context) []web.WebResponse {
	var responses []web.WebResponse

	// Mulai transaksi
	var err error
	tx := service.db.Begin()
	if tx.Error != nil {
		responses = append(responses, helper.WebResponsesFunc(http.StatusInternalServerError, "INTERNAL SERVER ERROR", "error starting transaction"))
		return responses
	}
	defer helper.CommitOrRollback(tx, &err)

	// Mengambil semua kursus dari repository
	courses, err := service.courseRepository.FindAll()
	if err != nil {
		responses = append(responses, helper.WebResponsesFunc(http.StatusInternalServerError, "INTERNAL SERVER ERROR", err.Error()))
		return responses
	}

	// Kembalikan response dengan data kursus
	for _, course := range courses {
		responses = append(responses, helper.WebResponsesFunc(http.StatusOK, "OK", course))
	}

	return responses

}

func (service CourseServiceImpl) Update(ctx *gin.Context, courseId int) web.WebResponse {
	//TODO implement me

	/// PUNYA AKUN USER
	/// ROLES TEACHER OR ADMIN
	/// FIND BY ID
	/// UPDATE

	panic("implement me")
}
