package schema

import (
	"gorm.io/gorm"
	"time"
)

type MedicineTransaction struct {
	ID                uint `gorm:"primarykey"`
	UserID            uint
	Name              string            `gorm:"not null"`
	Address           string            `gorm:"not null"`
	HP                string            `gorm:"not null"`
	PaymentMethod     string            `gorm:"type:enum('manual transfer bca', 'manual transfer bri', 'manual transfer bni')"`
	MedicineDetails   []MedicineDetails `gorm:"ForeignKey:MedicineTransactionID;references:ID"`
	TotalPrice        int
	StatusTransaction string `gorm:"type:enum('belum dibayar', 'sudah dibayar');default:'belum dibayar'"`
	UpdatedAt         time.Time
	CreatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
