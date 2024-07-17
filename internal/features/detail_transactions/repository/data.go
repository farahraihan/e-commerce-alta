package repository

import (
	"TokoGadget/internal/features/products"

	"gorm.io/gorm"
)

type DetailTransaction struct {
	gorm.Model
	ProductID     uint
	TransactionID uint
	Quantity      uint
	Product       products.Product `gorm:"foreignKey:ProductID"`
}
