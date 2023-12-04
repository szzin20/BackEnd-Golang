package schema

import (
	"gorm.io/gorm"
	"time"
)

type MedicineDetails struct {
	MedicineTransactionID uint     `gorm:"not null"`
	MedicineID            uint     `gorm:"not null"`
	Medicine              Medicine `gorm:"ForeignKey:MedicineID"`
	Quantity              int      `gorm:"not null"`
	TotalPriceMedicine    int      `gorm:"not null"`
}

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
