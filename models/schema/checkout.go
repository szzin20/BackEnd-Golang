package schema

import (
	"gorm.io/gorm"
	"time"
)

type Checkout struct {
	ID                    uint                `gorm:"primarykey"`
	MedicineTransactionID uint                `gorm:"not null"`
	PaymentConfirmation   string              `gorm:"not null"`
	PaymentStatus         string              `gorm:"type:enum('pending', 'success', 'cancelled');default:'pending'"`
	MedicineTransaction   MedicineTransaction `gorm:"ForeignKey:MedicineTransactionID;references:ID"`
	UpdatedAt             time.Time
	CreatedAt             time.Time
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
