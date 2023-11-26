package schema

import (
	"gorm.io/gorm"
	"time"
)

type MedicineTransaction struct {
	ID              uint `gorm:"primarykey"`
	UserID          uint
	Name            string `gorm:"not null"`
	Address         string `gorm:"not null"`
	HP              string `gorm:"not null"`
	PaymentMethod   string `gorm:"type:enum('manual transfer bca', 'manual transfer bri', 'manual transfer bni')"`
	TotalPrice      int
	UpdatedAt       time.Time
	CreatedAt       time.Time
	DeletedAt       gorm.DeletedAt    `gorm:"index"`
	MedicineDetails []MedicineDetails `gorm:"ForeignKey:MedicineTransactionID;references:ID"`
	Checkout        Checkout          `gorm:"ForeignKey:MedicineTransactionID;references:ID"`
}
