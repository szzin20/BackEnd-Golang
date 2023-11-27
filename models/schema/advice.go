package schema

import "time"

type Advice struct {
	ID            uint `gorm:"primaryKey"`
	TransactionID uint
	Message       string
	Image         string
	Audio         string
	CreatedAt     time.Time
	// ComplaintID uint
}
