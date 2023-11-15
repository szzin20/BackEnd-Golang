package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	Fullname  string
	Email     string `gorm:"not null;unique"`
	Password  string `gorm:"not null"`
	ImageURL  string
	Gender    string 
	Birthdate string
	BloodType string 
	Height    int
	Weight    int
	Role      string `gorm:"type:enum('user');default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
