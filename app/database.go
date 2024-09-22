package app

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jamal/tgs/helper"
	"jamal/tgs/models/domain"
)

// Global variable for the database connection

var DB *gorm.DB

// ConnectDatabase initializes the database connection and returns it
func ConnectDatabase() *gorm.DB {
	var err error
	DB, err = gorm.Open(mysql.Open("root:korie123@tcp(localhost:3306)/simple_lms"), &gorm.Config{})
	if err != nil {
		panic(errors.New("failed to connect to the database: " + err.Error()))
	}

	err = DB.AutoMigrate(
		&domain.User{},
		&domain.Course{},
		&domain.CourseContent{},
		&domain.CourseMember{},
		&domain.Completion{},
		&domain.Comment{},
	)
	helper.IfErrPanic(err)

	return DB
}
