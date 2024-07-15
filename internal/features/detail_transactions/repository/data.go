package repository

import (
	"TokoGadget/internal/features/products"

	"gorm.io/gorm"
)

type DetailTransaction struct {
	gorm.Model
	ProductID     uint             `json:"product_id"`
	Product       products.Product `gorm:"foreignKey:ProductID"`
	TransactionID uint             `json:"transaction_id"`
	Quantity      uint             `json:"quantity"`
}
