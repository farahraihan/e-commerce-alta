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

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Update(c echo.Context) error
	GetProfile(c echo.Context) error
	Delete(c echo.Context) error 
}

type UserServices interface {
	Register(newUser User) error
	Login(email string, password string) (User, string, error)
	UpdateProfile(userid uint, accounts User) error
	GetProfile(userid uint) (*User, error)
	DeleteAccount(userid uint) error
}

type UserQuery interface {
	Register(newUser User) error
	Login(email string) (User, error)
	UpdateAccount(userID uint, account User) error
	GetAccountByID(userid uint) (*User, error)	
	DeleteAccount(userid uint) error
}

type LoginValidate struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,alphanum"`
}
