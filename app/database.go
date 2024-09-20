package app

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jamal/tgs/helper"
	"jamal/tgs/models/domain"
)

// <-- usage for make connection to database

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:korie123@tcp(localhost:3306)/simple_lms"))
	if err != nil {
		panic(errors.New("failed connected data base"))
	}

	err = db.AutoMigrate(
		&domain.User{},
		&domain.Course{},
		&domain.CourseContent{},
		&domain.CourseMember{},
		&domain.Completion{},
		&domain.Comment{},
	)
	helper.IfErrPanic(err)

	DB = db

}
