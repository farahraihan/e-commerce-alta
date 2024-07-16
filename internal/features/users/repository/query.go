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

func (um *UserModel) UpdateAccount(userID uint, account users.User) error {
	var user User
	tx := um.db.First(&user, userID)
	if tx.Error != nil {
		return tx.Error
	}
	user.ProfilePicture = account.ProfilePicture
	user.Fullname = account.Fullname
	user.Email = account.Email
	user.Password = account.Password
	user.PhoneNumber = account.PhoneNumber
	user.Address = account.Address

	tx = um.db.Save(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (um *UserModel) GetAccountByID(userid uint) (*users.User, error) {
	var userData User
	tx := um.db.First(&userData, userid)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var user = users.User{
		ProfilePicture: userData.ProfilePicture,
		Fullname:       userData.Fullname,
		Email:          userData.Email,
		PhoneNumber:    userData.PhoneNumber,
		Address:        userData.Address,
	}

	return &user, nil
}

func (um *UserModel) DeleteAccount(userid uint) error {
	tx := um.db.Delete(&User{}, userid)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
