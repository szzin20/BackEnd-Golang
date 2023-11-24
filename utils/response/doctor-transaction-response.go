package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateDoctorTransactionResponse(doctorTransaction *schema.DoctorTransaction, doctor schema.Doctor) web.CreateDoctorTransactionResponse {
	return web.CreateDoctorTransactionResponse{
		ID:                  doctorTransaction.ID,
		Fullname:            doctor.Fullname,
		Specialist:          doctor.Specialist,
		Price:               doctor.Price,
		PaymentMethod:       doctorTransaction.PaymentMethod,
		PaymentStatus:       doctorTransaction.PaymentStatus,
		PaymentConfirmation: doctorTransaction.PaymentConfirmation,
	}
}

func ConvertToGetDoctorTransactionResponse(doctorTransaction schema.DoctorTransaction, doctor schema.Doctor) web.CreateDoctorTransactionResponse {
	return web.CreateDoctorTransactionResponse{
		ID:                  doctorTransaction.ID,
		Fullname:            doctor.Fullname,
		Specialist:          doctor.Specialist,
		Price:               doctor.Price,
		PaymentMethod:       doctorTransaction.PaymentMethod,
		PaymentStatus:       doctorTransaction.PaymentStatus,
		PaymentConfirmation: doctorTransaction.PaymentConfirmation,
	}
}

func ConvertToGetAllDoctorTransactionsResponse(doctorTransaction schema.DoctorTransaction, doctor schema.Doctor) web.DoctorTransactionsResponse {
	return web.DoctorTransactionsResponse{
		ID:       doctorTransaction.ID,
		Fullname: doctor.Fullname,
	}
}
