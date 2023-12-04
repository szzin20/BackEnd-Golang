package schema

import (
	"time"

	"gorm.io/gorm"
)

type DoctorTransaction struct {
	ID                  uint   `gorm:"primaryKey"`
	DoctorID            uint   `gorm:"foreignKey:DoctorID"`
	UserID              uint   `gorm:"foreignKey:UserID"`
	HealthDetails       string `gorm:"not null"`
	Price               int    `gorm:"not null"`
	PaymentMethod       string `gorm:"type:enum('manual transfer bca', 'manual transfer bri', 'manual transfer bni');default:null"`
	PaymentConfirmation string `gorm:"not null"`
	PaymentStatus       string `gorm:"type:enum('pending', 'success', 'cancelled');default:'pending'"`
	PatientStatus       string `gorm:"type:enum('pending', 'recovered', 'ongoing consultation', 'referred');default:'pending'"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *gorm.DeletedAt `gorm:"index"`
	Roomchat            Roomchat        `gorm:"ForeignKey:TransactionID;references:ID"` // one to one

}
