package domain

import "time"

type Completion struct {
	ID         uint `gorm:"primaryKey"`
	MemberID   uint
	ContentID  uint
	LastUpdate time.Time
}
