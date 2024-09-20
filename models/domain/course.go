package domain

type Course struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(200);not null"`
	Description string `gorm:"type:text"`
	Price       int
	Image       string `gorm:"type:varchar(200)"`
	Teacher     int

	/// Relasi ke CourseContent dan CourseMember
	CourseContents []CourseContent `gorm:"foreignKey:CourseID"`
	CourseMembers  []CourseMember  `gorm:"foreignKey:CourseID"`
}
