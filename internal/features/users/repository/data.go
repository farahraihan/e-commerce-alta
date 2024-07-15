package repository

import (
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/sales"
	"TokoGadget/internal/features/transactions"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname       string                     `json:"fullname"`
	Password       string                     `json:"password"`
	Email          string                     `json:"email"`
	PhoneNumber    string                     `json:"phone"`
	Address        string                     `json:"address"`
	ProfilePicture string                     `json:"profile_picture"`
	products       []products.Product         `gorm:"foreignKey:UserID"`
	transactions   []transactions.Transaction `gorm:"foreignKey:UserID"`
	sellers        []sales.Sales              `gorm:"foreignKey:SellerID"`
	buyers         []sales.Sales              `gorm:"foreignKey:BuyerID"`
}

// func (u *User) toUserEntity() users.User {
// 	return users.User{
// 		ID:             u.ID,
// 		Fullname:       u.Fullname,
// 		Password:       u.Password,
// 		Email:          u.Email,
// 		PhoneNumber:    u.PhoneNumber,
// 		Address:        u.Address,
// 		ProfilePicture: u.ProfilePicture,
// 	}
// }
