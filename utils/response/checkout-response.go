package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToGetCheckoutResponse(checkout *schema.Checkout) web.CheckoutResponse {
	medicineDetailsResponse := make([]web.MedicineDetailsResponse, len(checkout.MedicineTransaction.MedicineDetails))

	for i, md := range checkout.MedicineTransaction.MedicineDetails {
		medicineDetailsResponse[i] = web.MedicineDetailsResponse{
			MedicineID:         md.MedicineID,
			Quantity:           md.Quantity,
			TotalPriceMedicine: md.TotalPriceMedicine,
		}
	}

	MedicineCheckoutResponse := web.MedicineCheckoutResponse{
		UserID:                  checkout.MedicineTransaction.UserID,
		Name:                    checkout.MedicineTransaction.Name,
		Address:                 checkout.MedicineTransaction.Address,
		HP:                      checkout.MedicineTransaction.HP,
		PaymentMethod:           checkout.MedicineTransaction.PaymentMethod,
		MedicineDetailsResponse: medicineDetailsResponse,
		TotalPrice:              checkout.MedicineTransaction.TotalPrice, // Assuming MedicineTransaction has a TotalPrice field
		StatusTransaction:       checkout.MedicineTransaction.StatusTransaction,
	}

	return web.CheckoutResponse{
		ID:                       checkout.ID,
		PaymentStatus:            checkout.PaymentStatus,
		MedicineTransactionID:    checkout.MedicineTransactionID,
		MedicineCheckoutResponse: MedicineCheckoutResponse,
		CreatedAt:                checkout.CreatedAt,
		PaymentConfirmation:      checkout.PaymentConfirmation,
	}
}

func ConvertToGetAllCheckoutResponse(checkouts []schema.Checkout) []web.CheckoutResponse {
	var responses []web.CheckoutResponse

	for _, checkout := range checkouts {
		response := ConvertToGetCheckoutResponse(&checkout)
		responses = append(responses, response)
	}

	return responses
}
