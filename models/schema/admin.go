package schema

import (
	"time"
)

type Admin struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"type:enum('admin');default:'admin'"`
	UpdatedAt time.Time
}
