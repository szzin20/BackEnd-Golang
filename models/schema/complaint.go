package schema

import (
	"time"
)

type Complaint struct {
	ID                  uint `gorm:"primaryKey"`
	DoctorTransactionID uint `gorm:"foreignKey"`
	Title               string
	Content             string
	Status              ComplaintStatus `gorm:"type:enum('', '', '')"`
	CreatedAt           time.Time
	
}

type ComplaintStatus string
