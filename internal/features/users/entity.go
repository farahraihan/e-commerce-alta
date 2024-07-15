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
	Login() echo.HandlerFunc
}

type Services interface {
	Register(newUser User) error
	Login(email string, password string) (User, string, error)
}

type Query interface {
	Register(newUser User) error
	Login(email string) (User, error)
}

type LoginValidate struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,alphanum"`
}
