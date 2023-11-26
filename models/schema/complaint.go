package schema

import "time"

type Complaint struct {
	ID            uint `gorm:"primaryKey"`
	Message       string
	Image         string
	Audio         string
	CreatedAt     time.Time
	TransactionID uint
	Advice        Advice `gorm:"ForeignKey:ComplaintID;references:ID"` // one to one
}
