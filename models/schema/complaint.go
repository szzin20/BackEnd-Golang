package schema

import "time"

type Complaint struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Content   string
	Status    string
	CreatedAt time.Time
	TransactionID int
	Advice Advice `gorm:"ForeignKey:ComplaintID;references:ID"` // one to one
}

