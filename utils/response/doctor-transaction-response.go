package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToCreateDTResponse(DoctorTransaction *schema.DoctorTransaction) web.CreateDoctorTransactionResponse {
	return web.CreateDoctorTransactionResponse{
		DoctorID:      DoctorTransaction.DoctorID,
		UserID:        DoctorTransaction.UserID,
		HealthDetails: DoctorTransaction.HealthDetails,
		PaymentMethod: DoctorTransaction.PaymentMethod,
		Price:         DoctorTransaction.Price,
		ImageURL:      DoctorTransaction.ImageURL,
		PaymentStatus: DoctorTransaction.PaymentStatus,
	}
}

func ConvertToGetAllDTResponse(doctorTransactions []schema.DoctorTransaction) []web.CreateDoctorTransactionResponse {
	var results []web.CreateDoctorTransactionResponse
	for _, dt := range doctorTransactions {
		dtResponse := web.CreateDoctorTransactionResponse{
			DoctorID:      dt.DoctorID,
			UserID:        dt.UserID,
			HealthDetails: dt.HealthDetails,
			PaymentMethod: dt.PaymentMethod,
			Price:         dt.Price,
			ImageURL:      dt.ImageURL,
			PaymentStatus: dt.PaymentStatus,
		}
		results = append(results, dtResponse)
	}
	return results
}
