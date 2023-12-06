package response

import (
	"healthcare/models/schema"
	"healthcare/models/web"
	"strings"
)

func ConvertToUserRegisterResponse(user *schema.User) web.UserRegisterResponse {
	return web.UserRegisterResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
	}
}

func ConvertToUserLoginResponse(user schema.User) web.UserLoginResponse {
	return web.UserLoginResponse{
		Fullname: user.Fullname,
		Email:    user.Email,
	}
}

func ConvertToUserUpdateResponse(user *schema.User) web.UserUpdateResponse {
	bloodType := strings.ToUpper(user.BloodType)
	gender := strings.ToLower(user.Gender)
	return web.UserUpdateResponse{
		Fullname:       user.Fullname,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		Gender:         gender,
		Birthdate:      user.Birthdate,
		BloodType:      bloodType,
		Height:         user.Height,
		Weight:         user.Weight,
	}
}

func ConvertToGetUserResponse(user *schema.User) web.UserUpdateResponse {
	return web.UserUpdateResponse{
		Fullname:       user.Fullname,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		Gender:         user.Gender,
		Birthdate:      user.Birthdate,
		BloodType:      user.BloodType,
		Height:         user.Height,
		Weight:         user.Weight,
	}
}
func ConvertToGetUserIDbyAdminResponse(user *schema.User) web.UserAllResponseByAdmin {
	return web.UserAllResponseByAdmin{
		ID:             user.ID,
		Fullname:       user.Fullname,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		Gender:         user.Gender,
		Birthdate:      user.Birthdate,
		BloodType:      user.BloodType,
		Height:         user.Height,
		Weight:         user.Weight,
	}
}

func ConvertToGetAllUsersResponse(users []schema.User) []web.UserUpdateResponse {
	var results []web.UserUpdateResponse
	for _, user := range users {
		userResponse := web.UserUpdateResponse{
			Fullname:       user.Fullname,
			Email:          user.Email,
			ProfilePicture: user.ProfilePicture,
			Gender:         user.Gender,
			Birthdate:      user.Birthdate,
			BloodType:      user.BloodType,
			Height:         user.Height,
			Weight:         user.Weight,
		}
		results = append(results, userResponse)
	}
	return results
}

func ConvertToGetAllUserByAdminResponse(users []schema.User) []web.UserAllResponseByAdmin {
	var results []web.UserAllResponseByAdmin

	// Iterasi melalui setiap dokter dan konversi ke format respons
	for _, user := range users {
		userResponse := web.UserAllResponseByAdmin{
			ID:             user.ID,
			Fullname:       user.Fullname,
			Email:          user.Email,
			ProfilePicture: user.ProfilePicture,
			Gender:         user.Gender,
			Birthdate:      user.Birthdate,
			BloodType:      user.BloodType,
			Height:         user.Height,
			Weight:         user.Weight,
		}

		results = append(results, userResponse)
	}

	return results
}

func ConvertToGetUserbyAdmin(user schema.User) web.UserGetResponsebyAdmin {
    response := web.UserGetResponsebyAdmin{
        ID:         user.ID,
		ProfilePicture: user.ProfilePicture,
        Fullname:       user.Fullname,
        Email:  user.Email,
		Gender:         user.Gender,
		Birthdate:      user.Birthdate,
		BloodType:      user.BloodType,
		Height:         user.Height,
		Weight:         user.Weight,

    }

    doctorTransaction := ListConvertToGetUserTransactionbyAdminResponse(user.DoctorTransaction)
    response.DoctorTransaction = doctorTransaction
    medicineTransaction := ListConvertToGetUserMedicineTransactionbyAdminResponse(user.MedicineTransaction)
    response.MedicineTransaction = medicineTransaction

    return response
}