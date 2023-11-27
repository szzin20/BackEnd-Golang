package schema

import "time"

type Article struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null;unique"`
	Content   string `gorm:"not null"`
	Image     string `gorm:"default:null"`
	CreatedAt time.Time
	DoctorID  uint
}
