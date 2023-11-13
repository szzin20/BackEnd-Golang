package web

type AdminLoginResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
type AdminUpdateResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
