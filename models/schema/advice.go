package schema

import "time"

type Advice struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Content     string
	Status      string
	CreatedAt   time.Time
	ComplaintID int
}
