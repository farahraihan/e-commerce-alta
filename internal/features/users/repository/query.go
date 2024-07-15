package repository

import (
	"TokoGadget/internal/features/users"

	"gorm.io/gorm"
)

type UserModel struct {
	db *gorm.DB
}

func NewUserModel(connection *gorm.DB) users.Query {
	return &UserModel{
		db: connection,
	}
}

func (um *UserModel) Register(newUser users.User) error {
	cnv := toUserRegis(newUser)
	err := um.db.Create(&cnv).Error
	// err := um.db.Create(&newUser).Error
	return err
}
func (um *UserModel) Login(email string) (users.User, error) {
	var result User
	err := um.db.Where("email = ?", email).First(&result).Error
	if err != nil {
		return users.User{}, err
	}
	return result.toUserEntity(), nil
}
