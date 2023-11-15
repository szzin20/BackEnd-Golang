package schema

import "time"

type Advice struct {
	ID          uint `gorm:"primaryKey"`
	Message     string
	Image       string
	Audio       string
	CreatedAt   time.Time
	ComplaintID uint
}
