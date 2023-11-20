package schema

import (
	"time"
)

type DoctorTransaction struct {
	ID            uint   `gorm:"primaryKey"`
	DoctorID      uint   `gorm:"foreignKey:DoctorID"`
	UserID        uint   `gorm:"foreignKey:UserID"`
	HealthDetails string `gorm:"not null"`
	PaymentMethod string `gorm:"not null"`
	Price         int    `gorm:"not null"`
	ImageURL      string `gorm:"not null"`
	//Complaint     Complaint `gorm:"ForeignKey:TransactionID;references:ID"` // one to one
	PaymentStatus string `gorm:"type:enum('pending', 'success', 'cancelled');default:'pending'"`
	CreatedAt     time.Time
}
