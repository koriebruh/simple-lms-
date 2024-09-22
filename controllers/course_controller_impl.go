package controllers

import (
	"github.com/gin-gonic/gin"
	"jamal/tgs/services"
	"net/http"
	"strconv"
)

type CourseControllerImpl struct {
	courseService services.CourseService
}

func NewCourseController(service services.CourseService) CourseController {
	return &CourseControllerImpl{
		courseService: service,
	}
}

func (controller CourseControllerImpl) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller CourseControllerImpl) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller CourseControllerImpl) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller CourseControllerImpl) FindById(c *gin.Context) {
	// Ambil ID dari parameter URL
	courseIdParam := c.Param("id")
	courseId, err := strconv.Atoi(courseIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "BAD REQUEST",
			"message": "Invalid course ID",
		})
		return
	}

	// Panggil service untuk mencari kursus
	response := controller.courseService.FindById(c, courseId)

	// Kembalikan response ke klien
	c.JSON(response.Code, response)
}

func (controller CourseControllerImpl) FindAll(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
