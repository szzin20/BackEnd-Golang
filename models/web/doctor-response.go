package web

import "healthcare/models/schema"

// type DoctorResgisterResponse struct {
// 	Fullname           string `json:"fullname" form:"fullname"`
// 	Email              string `json:"email" form:"email"`
// 	Price              int    `json:"price" form:"price"`
// 	Tag                string `json:"tag" form:"tag"`
// 	ProfilePicture     string `json:"profile_picture" form:"profile_picture"`
// 	RegistrationLetter int    `json:"registration_letter" form:"registration_letter"`
// }

type DoctorLoginResponse struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}

type DoctorLogOutResponse struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Tag      string `json:"tag" form:"tag"`
}

type DoctorUpdateResponse struct {
	Fullname                string `json:"fullname" form:"fullname"`
	Email                   string `json:"email" form:"email"`
	Price                   int    `json:"price" form:"price"`
	Tag                     string `json:"tag" form:"tag"`
	ProfilePicture          string `json:"profile_picture" form:"profile_picture"`
	RegistrationCertificate int    `json:"registration_certificate" form:"registration_certificate"`
}

type DoctorAllResponse struct {
	Fullname                string              `json:"fullname" form:"fullname"`
	Email                   string              `json:"email" form:"email"`
	Price                   int                 `json:"price" form:"price"`
	Tag                     string              `json:"tag" form:"tag"`
	Status                  schema.DoctorStatus `json:"status" form:"status"`
	ProfilePicture          string              `json:"profile_picture" form:"profile_picture"`
	RegistrationCertificate int                 `json:"registration_certificate" form:"registration_certificate"`
}

// type DoctorPatientsResponse struct {
// 	DoctorID       uint                  `json:"doctor_id"`
// 	TransactionID  uint                  `json:"transaction_id"`
// 	DoctorFullname string                `json:"doctor_fullname"`
// 	Patients       []PatientInfoResponse `json:"patients"`
// }

// type PatientInfoResponse struct {
// 	ID        uint   `json:"id"`
// 	Name      string `json:"name"`
// 	Email     string `json:"email"`
// 	Gender    string `json:"gender"`
// 	Status    string `json:"status"`
// 	Height    int    `json:"height"`
// 	Weight    int    `json:"weight"`
// 	Birthdate string `json:"birthdate"`
// 	BloodType string `json:"blood_type"`
// }
