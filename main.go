package main

import (
	"github.com/gin-gonic/gin"
	"jamal/tgs/app"
	"jamal/tgs/controllers"
	"jamal/tgs/repository"
	"jamal/tgs/services"
	"log"
)

func main() {

	// Connect to database
	db := app.ConnectDatabase()

	// Initialize repositories
	courseRepo := repository.NewCourseRepository(db)

	// Initialize services
	courseService := services.NewCourseServiceImpl(courseRepo, db)

	// Initialize controllers
	courseController := controllers.NewCourseController(courseService)

	// Set up Gin router
	router := gin.Default()

	// Define routes
	router.POST("/courses", courseController.Create)
	router.GET("/courses", courseController.FindAll)
	router.GET("/courses/:id", courseController.FindById)
	router.PUT("/courses/:id", courseController.Update)
	router.DELETE("/courses/:id", courseController.Delete)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
