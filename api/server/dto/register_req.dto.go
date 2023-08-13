package dto

type RegisterRequestDTO struct {
	Email    string `json:"email" form:"email"`
	FullName string `json:"fullname" form:"fullname"`
	Password string `json:"password" form:"password"`
}
