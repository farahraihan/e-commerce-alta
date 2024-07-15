package handler

import (
	"TokoGadget/internal/features/users"
)
type LoginResponse struct {
	Token    string `json:"token"`
}

func ToLoginReponse(input users.User, tkn string) LoginResponse {
	return LoginResponse{
		Token:    tkn,
	}
}