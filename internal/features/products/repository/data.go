package repository

import (
	detailtransactions "TokoGadget/internal/features/detail_transactions"
	"TokoGadget/internal/features/products"
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
	Price              uint64                                 `json:"price"`
	Stock              int32                                  `json:"stock"`
	ProductPicture     string                                 `json:"product_picture"`
	DeletedAt          gorm.DeletedAt                         `gorm:"index"`
	detailtransactions []detailtransactions.DetailTransaction `gorm:"foreignKey:ProductID"`
	sales              []sales.Sales                          `gorm:"foreignKey:ProductID"`
}

func (p *Product) ToProductEntity() products.Product {
	return products.Product{
		ID:             p.ID,
		UserID:         p.UserID,
		ProductName:    p.ProductName,
		Category:       p.Category,
		Description:    p.Description,
		Price:          p.Price,
		Stock:          p.Stock,
		ProductPicture: p.ProductPicture,
	}
}

func ToProductData(input products.Product) Product {
	return Product{
		UserID:         input.UserID,
		ProductName:    input.ProductName,
		Category:       input.Category,
		Description:    input.Description,
		Price:          input.Price,
		Stock:          input.Stock,
		ProductPicture: input.ProductPicture,
	}
}
