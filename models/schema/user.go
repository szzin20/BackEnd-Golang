package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int `gorm:"primarykey"`
	Fullname  string
	Email     string
	Password  string
	ImageURL  string
	Gender    string `gorm:"type:enum('male', 'female')"`
	Birthdate string
	BloodType string `gorm:"type:enum('A', 'B', 'O', 'AB')"`
	Height    int
	Weight    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
