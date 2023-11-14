package web

type AdminLoginRequest struct {
	Email    string `json:"email" form:"email" `
	Password string `json:"password" form:"password" `
}

type AdminUpdateRequest struct {
	Name     string `json:"name" form:"name" `
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password" `
}

type AdminResetPasswordRequest struct {
	Email              string `json:"email" form:"email" `
	NewPassword        string `json:"new_password" form:"new_password" `
	ConfirmNewPassword string `json:"confirm_new_password" form:"confirm_new_password" `
}