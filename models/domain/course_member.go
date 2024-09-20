package domain

type CourseMember struct {
	ID       uint `gorm:"primaryKey"`
	CourseID uint
	UserID   uint
	Role     string `gorm:"type:enum('student', 'teacher', 'admin')"`
}
