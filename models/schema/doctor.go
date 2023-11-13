package schema

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID                 int    `gorm:"primarykey"`
	Fullname           string `gorm:"not null"`
	Email              string `gorm:"not null;unique"`
	Password           string `gorm:"not null"`
	Status             bool   
	Price              int    `gorm:"not null"`
	Tag                string `gorm:"not null"`
	ImageURL           string `gorm:"not null"`
	RegistrationLetter string `gorm:"not null"`
	Role               string `gorm:"type:enum('doctor');default:'doctor'"`
	UpdatedAt          time.Time
	CreatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
