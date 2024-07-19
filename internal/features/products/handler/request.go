package handler

import "TokoGadget/internal/features/products"

type CreateOrUpdateProductRequest struct {
	ProductName    string `form:"product_name"`
	Category       string `form:"category"`
	Description    string `form:"description"`
	Price          uint64 `form:"price"`
	Stock          int32  `form:"stock"`
	ProductPicture string `form:"product_picture"`
}

func ToModelProduct(r CreateOrUpdateProductRequest, userID uint) products.Product {
	return products.Product{
		ProductName:    r.ProductName,
		Category:       r.Category,
		Description:    r.Description,
		Price:          r.Price,
		Stock:          r.Stock,
		ProductPicture: r.ProductPicture,
		UserID:         userID,
	}
}
