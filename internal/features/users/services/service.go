package services

import (
	"TokoGadget/internal/features/users"
	"TokoGadget/internal/utils"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
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