package repository

import (
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/sales"

	"gorm.io/gorm"
)

type DetailTransaction struct {
	gorm.Model
	ProductID     uint
	TransactionID uint
	Quantity      uint
	Product       products.Product `gorm:"foreignKey:ProductID"`
	sales         []sales.Sale     `gorm:"foreignKey:TransactionID"`
}
