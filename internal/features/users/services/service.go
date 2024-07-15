package services

import (
	"TokoGadget/internal/features/users"
	"TokoGadget/internal/utils"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServices struct {
	qry      users.Query
	validate *validator.Validate
}

func NewUserService(q users.Query) users.Services {
	return &UserServices{
		qry:      q,
		validate: validator.New(),
	}
}

func (us *UserServices) Register(newData users.User) error {
	processPw, err := utils.GeneratePassword(newData.Password)

	if err != nil {
		log.Println("register generate password error:", err.Error())
		return errors.New("input data tidak valid, data tidak bisa diproses")
	}

	newData.Password = string(processPw)

	err = us.qry.Register(newData)

	if err != nil {
		log.Println("register sql error:", err.Error())
		return errors.New("terjadi kesalahan pada server saat mengolah data")
	}

	return nil
}

func (us *UserServices) Login(email string, password string) (users.User, string, error) {
	err := us.validate.Struct(&users.LoginValidate{Email: email, Password: password})
	msg := "terjadi kesalahan pada server"

	if err != nil {
		log.Println("login validation error", err.Error())
		return users.User{}, "", errors.New("validasi tidak sesuai")
	}

	result, err := us.qry.Login(email)
	if err != nil {
		log.Println("login sql error:", err.Error())
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			msg = "data tidak ditemukan"
		}
		return users.User{}, "", errors.New(msg)
	}

	err = utils.CheckPassword([]byte(password), []byte(result.Password))
	if err != nil {
		log.Println("login hash error:", err.Error())
		if err.Error() == bcrypt.ErrMismatchedHashAndPassword.Error() {
			msg = "data tidak ditemukan"
		}
		return users.User{}, "", errors.New(msg)
	}

	token, err := utils.GenerateToken(result.ID)
	if err != nil {
		log.Println("login jwt error:", err.Error())
		if err.Error() == jwt.ErrTokenMalformed.Error() {
			msg = "data tidak dapat diproses"
		}
		return users.User{}, "", errors.New(msg)
	}

	return result, token, nil
}

