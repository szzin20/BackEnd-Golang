package schema

import (
	"time"

	"gorm.io/gorm"
)

type DoctorStatus string

const (
	Online  DoctorStatus = "Online"
	Offline DoctorStatus = "Offline"
)

type Doctor struct {
	ID                      uint                `gorm:"primarykey"`
	Fullname                string              `gorm:"not null"`
	Email                   string              `gorm:"not null;unique"`
	Password                string              `gorm:"not null"`
	Status                  DoctorStatus        `gorm:"type:enum('Online', 'Offline')"`
	Price                   int                 `gorm:"not null"`
	Tag                     string              `gorm:"not null"`
	ProfilePicture          string              `gorm:"not null"`
	RegistrationCertificate int                 `gorm:"not null"`
	Role                    string              `gorm:"type:enum('doctor');default:'doctor'"`
	DoctorTransaction       []DoctorTransaction `gorm:"ForeignKey:DoctorID;references:ID"`
	UpdatedAt               time.Time
	CreatedAt               time.Time
	DeletedAt               gorm.DeletedAt `gorm:"index"`
}
