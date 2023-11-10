package schema

import (
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	ID                 int    `gorm:"primarykey"`
	Name               string `gorm:"not null"`
	Email              string `gorm:"not null;unique"`
	Password           string `gorm:"not null"`
	Status             bool   `gorm:"not null"`
	Price              int    `gorm:"not null"`
	Tag                string `gorm:"not null"`
	ImageURL           string `gorm:"not null"`
	RegistrationLetter string `gorm:"not null"`
	UpdatedAt          time.Time
	CareateAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
