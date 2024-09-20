package domain

type CourseContent struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string `gorm:"type:varchar(200);not null"`
	Description    string `gorm:"type:text"`
	VideoURL       string `gorm:"type:varchar(255)"`
	FileAttachment string `gorm:"type:varchar(255)"`
	ParentID       uint
	CourseID       uint

	/// Relasi ke Completion dan Comment
	Completions []Completion `gorm:"foreignKey:ContentID"`
	Comments    []Comment    `gorm:"foreignKey:ContentID"`
}
