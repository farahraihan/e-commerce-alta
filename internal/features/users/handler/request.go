package handler

import (
	"TokoGadget/internal/features/users"
)

type RegisterRequest struct {
	Fullname string          `json:"fullname"`
	Password string          `json:"password"`
	Email    string          `json:"email"`
}

func ToModelUsers(r RegisterRequest) users.User {
	return users.User{
		Fullname: r.Fullname,
		Password: r.Password,
		Email:    r.Email,
	}
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
