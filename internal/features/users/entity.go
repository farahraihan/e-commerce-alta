package users

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	ID             uint
	Fullname       string
	Password       string
	Email          string
	PhoneNumber    string
	Address        string
	ProfilePicture string
}

type Handler interface {
	Register() echo.HandlerFunc
}

type Services interface {
	Register(newUser User) error
}

type Query interface {
	Register(newUser User) error
}
