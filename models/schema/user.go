package schema

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                  uint `gorm:"primarykey"`
	Fullname            string
	Email               string `gorm:"not null"`
	Password            string `gorm:"not null"`
	ProfilePicture      string
	Gender              string `gorm:"type:enum('male', 'female');default:null"`
	Birthdate           string
	BloodType           string `gorm:"type:enum('A', 'B', 'O', 'AB');default:null"`
	Height              int
	Weight              int
	OTP                 string `gorm:"not null"`
	Role                string `gorm:"type:enum('user');default:'user'"`
	IsVerified          bool   `gorm:"not null;default:false"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt        `gorm:"index"`
	MedicineTransaction []MedicineTransaction `gorm:"foreignKey:UserID;references:ID"`
	DoctorTransaction   []DoctorTransaction   `gorm:"foreignKey:UserID;references:ID"`
}
