package schema

import (
	// "time"

	"gorm.io/gorm"
)

type DoctorTransaction struct {
	// ID                  uint   `gorm:"primaryKey"`
	gorm.Model
	DoctorID            uint   `gorm:"foreignKey:DoctorID"`
	UserID              uint   `gorm:"foreignKey:UserID"`
	HealthDetails       string `gorm:"not null"`
	PaymentMethod       string `gorm:"type:enum('manual', 'gateway')"`
	Price               int    `gorm:"not null"`
	PaymentConfirmation string `gorm:"not null"`
	PaymentStatus       string `gorm:"type:enum('pending', 'success', 'cancelled');default:'pending'"`
	PatientStatus       string `gorm:"type:enum('pending', 'solved', 'unsolved');default:'pending'"`
	// CreatedAt           time.Time
	// UpdatedAt           time.Time
	// DeletedAt           *gorm.DeletedAt `gorm:"index"`
	Doctor              Doctor          //`gorm:"embedded"`
	// Complaint     Complaint `gorm:"ForeignKey:TransactionID;references:ID"` // one to one
}
