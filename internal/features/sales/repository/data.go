package repository

import (
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/users"

	"gorm.io/gorm"
)

type Sales struct {
	gorm.Model
	SellerID  uint             `json:"seller_id"`
	Seller    users.User       `gorm:"foreignKey:SellerID"`
	BuyerID   uint             `json:"buyer_id"`
	Buyer     users.User       `gorm:"foreignKey:BuyerID"`
	ProductID uint             `json:"product_id"`
	Product   products.Product `gorm:"foreignKey:ProductID"`
	Quantity  int              `json:"quantity"`
	Status    string           `json:"status"`
}
