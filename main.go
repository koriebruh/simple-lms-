package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jamal/tgs/app"
)

func main() {

	app.ConnectDatabase()
	fmt.Println("ykosooo")

	r := gin.Default()
	r.GET("api/courses")     // GET ALL COURSE
	r.GET("api/courses/:id") // GET COURSE BY ID
	r.GET("api/profile/:id")
}
