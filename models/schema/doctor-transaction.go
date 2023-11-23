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
	ConsultationTime    string `gorm:"type:enum('30 menit');default:'30 menit'"`
	Price               int    `gorm:"not null"`
	PaymentMethod       string `gorm:"type:enum('manual transfer bca', 'manual transfer bri', 'manual transfer bni')"`
	PaymentConfirmation string `gorm:"not null"`
	PaymentStatus       string `gorm:"type:enum('pending', 'success', 'cancelled');default:'pending'"`
	PatientStatus       string `gorm:"type:enum('pending', 'solved', 'unsolved');default:'pending'"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           *gorm.DeletedAt `gorm:"index"`
}
