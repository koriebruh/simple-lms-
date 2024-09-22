package domain

type CourseFindAll struct {
	CourseID    uint    `json:"course_id"`
	CourseName  string  `json:"course_name"`
	CoursePrice uint    `json:"course_price"`
	Teacher     Teacher `json:"teacher"`
}

type Teacher struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}
