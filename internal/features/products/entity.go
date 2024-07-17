package products

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Product struct {
	ID             uint
	UserID         uint
	ProductName    string
	Category       string
	Description    string
	Price          uint64
	Stock          int32
	ProductPicture string
	CreatedAt      time.Time      `gorm:"default:current_timestamp"`
	UpdatedAt      time.Time      `gorm:"default:current_timestamp"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type PHandler interface {
	AddProduct() echo.HandlerFunc
	GetAllProducts() echo.HandlerFunc
	GetProductByID() echo.HandlerFunc
	UpdateProductByID() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
}

type PServices interface {
	AddProduct(newProduct Product) error
	GetAllProducts() ([]Product, error)
	GetAllProductsByUserID(user_id uint) ([]Product, error)
	GetProductByID(id uint) (*Product, error)
	UpdateProductByID(id uint, updatedProduct Product) error
	DeleteProduct(id uint) error
	GetProductsBySearch(search string) ([]Product, error)
}

type PQuery interface {
	AddProduct(newProduct Product) error
	GetAllProducts() ([]Product, error)
	GetAllProductsByUserID(user_id uint) ([]Product, error)
	GetProductByID(id uint) (*Product, error)
	UpdateProductByID(id uint, updatedProduct Product) error
	DeleteProduct(id uint) error
	GetProductsBySearch(search string) ([]Product, error)
}
