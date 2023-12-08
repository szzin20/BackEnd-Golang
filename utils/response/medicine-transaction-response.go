package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToMedicineTransactionResponse(mt *schema.MedicineTransaction) *web.MedicineTransactionResponse {
	medicineDetailsResponse := make([]web.MedicineDetailsResponse, len(mt.MedicineDetails))

	for i, md := range mt.MedicineDetails {
		medicineDetailsResponse[i] = web.MedicineDetailsResponse{
			MedicineID:         md.MedicineID,
			Quantity:           md.Quantity,
			TotalPriceMedicine: md.TotalPriceMedicine,
		}
	}

	return &web.MedicineTransactionResponse{
		ID:                      mt.ID,
		UserID:                  mt.UserID,
		Name:                    mt.Name,
		Address:                 mt.Address,
		HP:                      mt.HP,
		PaymentMethod:           mt.PaymentMethod,
		MedicineDetailsResponse: medicineDetailsResponse,
		TotalPrice:              mt.TotalPrice,
		StatusTransaction:       mt.StatusTransaction,
		CreatedAt:               mt.CreatedAt,
	}
}

func ConvertToMedicineTransactionListResponse(medicineTransactions []schema.MedicineTransaction) []web.MedicineTransactionResponse {
	var results []web.MedicineTransactionResponse
	for _, mt := range medicineTransactions {
		medicineDetailsResponse := make([]web.MedicineDetailsResponse, len(mt.MedicineDetails))
		for i, md := range mt.MedicineDetails {
			medicineDetailsResponse[i] = web.MedicineDetailsResponse{
				MedicineID:         md.MedicineID,
				Quantity:           md.Quantity,
				TotalPriceMedicine: md.TotalPriceMedicine,
			}
		}
		medicineTransactionResponse := web.MedicineTransactionResponse{
			ID:                      mt.ID,
			UserID:                  mt.UserID,
			Name:                    mt.Name,
			Address:                 mt.Address,
			HP:                      mt.HP,
			PaymentMethod:           mt.PaymentMethod,
			MedicineDetailsResponse: medicineDetailsResponse,
			TotalPrice:              mt.TotalPrice,
			StatusTransaction:       mt.StatusTransaction,
			CreatedAt:               mt.CreatedAt,
		}
		results = append(results, medicineTransactionResponse)
	}
	return results
}