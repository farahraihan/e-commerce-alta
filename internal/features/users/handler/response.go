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

type UserResponse struct {
	Id				uint 	`json:"id"`
	ProfilePicture string `json:"profile_picture"`
	FullName       string `json:"fullname"`
	Email          string `json:"email"`
	PhoneNumber    string `json:"phone_number"`
	Address        string `json:"address"`
}

func ToGetUserResponse(user users.User) UserResponse {
	return UserResponse{
		Id:             user.ID,
		ProfilePicture: user.ProfilePicture,
		FullName:       user.Fullname,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		Address:        user.Address,
	}
}
