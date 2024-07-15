package repository

import (
	detailtransactions "TokoGadget/internal/features/detail_transactions"
	"TokoGadget/internal/features/sales"
	"TokoGadget/internal/features/users"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID             uint                                   `json:"json_id"`
	User               users.User                             `gorm:"foreignKey:UserID"`
	ProductName        string                                 `json:"product_name"`
	Category           string                                 `json:"category"`
	Description        string                                 `json:"description"`
	Price              int64                                  `json:"price"`
	Stock              int32                                  `json:"stock"`
	ProductPicture     string                                 `json:"product_picture"`
	detailtransactions []detailtransactions.DetailTransaction `gorm:"foreignKey:ProductID"`
	sales              []sales.Sales                          `gorm:"foreignKey:ProductID"`
}
