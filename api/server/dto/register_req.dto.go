package dto

type RegisterRequestDTO struct {
	Email    string `json:"email" form:"email" valid:"email"`
	FullName string `json:"fullname" form:"fullname" valid:"-"`
	Password string `json:"password" form:"password" valid:"-"`
}
