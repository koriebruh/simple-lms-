package domain

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(50);not null"`
	Email    string `gorm:"type:varchar(100);not null"`
	Password string `gorm:"type:varchar(100);not null"`
	FullName string `gorm:"type:varchar(100)"`

	/// Relasi ke CourseMember dan Comment
	CourseMembers []CourseMember `gorm:"foreignKey:UserID"`
	Comments      []Comment      `gorm:"foreignKey:UserID"`
}
