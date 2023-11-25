package schema

import "time"

type Complaint struct {
	ID            uint `gorm:"primaryKey"`
	TransactionID uint
	Message       string
	Image         string
	Audio         string
	CreatedAt     time.Time
	Advice        Advice `gorm:"ForeignKey:ComplaintID;references:ID"` // one to one
}
