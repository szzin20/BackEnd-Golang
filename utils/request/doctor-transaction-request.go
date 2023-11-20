package request

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateDTRequest(DoctorTransaction web.CreateDoctorTransactionRequest, userID, doctorID uint, fullname string, specialist string, price int) *schema.DoctorTransaction {
	return &schema.DoctorTransaction{
		UserID: userID,
		DoctorID: doctorID,
		PaymentMethod:       DoctorTransaction.PaymentMethod,
		PaymentConfirmation: DoctorTransaction.PaymentConfirmation,
	}
}
