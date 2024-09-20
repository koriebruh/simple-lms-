package domain

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	ContentID uint
	Comment   string `gorm:"type:text"`
}
