package handler

import (
	"TokoGadget/internal/features/products"
	"time"
)

type ProductResponse struct {
	ID             uint      `json:"id"`
	ProductName    string    `json:"product_name"`
	Category       string    `json:"category"`
	Description    string    `json:"description"`
	Price          uint64    `json:"price"`
	Stock          int32     `json:"stock"`
	ProductPicture string    `json:"product_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      *string   `json:"deleted_at,omitempty"`
}

func ToResponseProduct(p products.Product) ProductResponse {
	var deletedAt *string
	if p.DeletedAt.Valid {
		deletedAtValue := p.DeletedAt.Time.String()
		deletedAt = &deletedAtValue
	}
	return ProductResponse{
		ID:             p.ID,
		ProductName:    p.ProductName,
		Category:       p.Category,
		Description:    p.Description,
		Price:          p.Price,
		Stock:          p.Stock,
		ProductPicture: p.ProductPicture,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		DeletedAt:      deletedAt,
	}
}

func ToResponseProducts(products []products.Product) []ProductResponse {
	response := make([]ProductResponse, len(products))
	for i, p := range products {
		response[i] = ToResponseProduct(p)
	}
	return response
}
