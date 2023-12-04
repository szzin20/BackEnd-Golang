package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
)

func ConvertToAdminLoginResponse(admin schema.Admin) web.AdminLoginResponse {
	return web.AdminLoginResponse{
		Name:  admin.Name,
		Email: admin.Email,
	}
}
func ConvertToAdminUpdateResponse(admin *schema.Admin) web.AdminUpdateResponse {
	return web.AdminUpdateResponse{
		Name:  admin.Name,
		Email: admin.Email,
	}
}

func ConvertToGetProfileAdminResponse(admin *schema.Admin) web.AdminProfileResponse {
	return web.AdminProfileResponse{
		Name:  admin.Name,
		Email: admin.Email,
	}
}

func ConvertToAdminTransactionUsersResponse(transactions []schema.DoctorTransaction) []web.AdminTransactionUsersResponse {
	var results []web.AdminTransactionUsersResponse
	for _, transaction := range transactions {
		adminsResponse := web.AdminTransactionUsersResponse{
			TransactionID:       transaction.ID,
			DoctorID:            transaction.DoctorID,
			UserID:              transaction.UserID,
			PaymentMethod:       transaction.PaymentMethod,
			Price:               transaction.Price,
			CreatedAt:           transaction.CreatedAt,
			PaymentConfirmation: transaction.PaymentConfirmation,
			PaymentStatus:       transaction.PaymentStatus,
		}
		results = append(results, adminsResponse)
	}
	return results
}
