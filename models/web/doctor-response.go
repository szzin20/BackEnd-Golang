package web

type DoctorResgisterResponse struct {
	Fullname       string `json:"fullname" form:"fullname" `
	Email          string `json:"email" form:"email"`
	Price          int    `json:"price" form:"price" `
	Gender         string `json:"gender" form:"gender" `
	Specialist     string `json:"specialist" form:"specialist" `
	ProfilePicture string `json:"profile_picture" form:"profile_picture" `
	NoSTR          int    `json:"no_str" form:"no_str" `
	Experience     string `json:"experience" form:"experience" `
}

type DoctorLoginResponse struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Token    string `json:"token" form:"token"`
}

type DoctorUpdateResponse struct {
	Fullname       string `json:"fullname" form:"fullname" `
	Email          string `json:"email" form:"email"`
	Price          int    `json:"price" form:"price" `
	Gender         string `json:"gender" form:"gender" `
	Specialist     string `json:"tag" form:"tag" `
	ProfilePicture string `json:"profile_picture" form:"profile_picture" `
	NoSTR          int    `json:"registration_certificate" form:"registration_certificate" `
	Experience     string `json:"experience" form:"experience" `
}

type DoctorAllResponse struct {
	Fullname   string `json:"fullname" form:"fullname" `
	Price      int    `json:"price" form:"price" `
	Gender     string `json:"gender" form:"gender"`
	Status     bool   `json:"status" form:"status"`
	Specialist string `json:"tag" form:"tag" `
	NoSTR      int    `json:"registration_certificate" form:"registration_certificate" `
	Experience string `json:"experience" form:"experience" `
}

type DoctorResponse struct {
	Fullname       string `json:"fullname" form:"fullname" `
	Email          string `json:"email" form:"email"`
	Password       string `json:"password" form:"password" `
	Gender         string `json:"gender" form:"gender" `
	Specialist     string `json:"specialist" form:"specialist" `
	ProfilePicture string `json:"profile_picture" form:"profile_picture" `
	NoSTR          int    `json:"no_str" form:"no_str" `
	Experience     string `json:"experience" form:"experience" `
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
