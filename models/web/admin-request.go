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
	NewPassword        string `json:"newPassword" form:"password" `
	ConfirmNewPassword string `json:"confirmNewPassword" form:"confirmNewPassword" `
}
