package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateDTResponse(doctorTransaction *schema.DoctorTransaction, doctor schema.Doctor) web.CreateDoctorTransactionResponse {
	return web.CreateDoctorTransactionResponse{
		Fullname:            doctor.Fullname,
		Specialist:          doctor.Specialist,
		Price:               doctor.Price,
		PaymentMethod:       doctorTransaction.PaymentMethod,
		PaymentConfirmation: doctorTransaction.PaymentConfirmation,
		PaymentStatus:       doctorTransaction.PaymentStatus,
	}
}

func ConvertToGetAllDTResponse(doctorTransaction schema.DoctorTransaction, doctor schema.Doctor ) web.CreateDoctorTransactionResponse {
	return web.CreateDoctorTransactionResponse{
		Fullname:            doctor.Fullname,
		Specialist:          doctor.Specialist,
		Price:               doctor.Price,
		PaymentMethod:       doctorTransaction.PaymentMethod,
		PaymentConfirmation: doctorTransaction.PaymentConfirmation,
		PaymentStatus:       doctorTransaction.PaymentStatus,
	}
}
