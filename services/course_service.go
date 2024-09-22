package services

import (
	"github.com/gin-gonic/gin"
	"jamal/tgs/models/web"
)

type CourseService interface {
	Create(ctx *gin.Context) web.CourseCreateRequest
	Delete(ctx *gin.Context, courseId int) web.WebResponse
	FindById(ctx *gin.Context, courseId int) web.WebResponse
	FindAll(ctx *gin.Context) []web.WebResponse
	Update(ctx *gin.Context, courseId int) web.WebResponse
}
