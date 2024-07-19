package products

import (
	"TokoGadget/internal/features/users"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Product struct {
	ID             uint
	UserID         uint
	User           users.User
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
	GetAllProducts(term string, limit int, offset int) ([]Product, error)
	GetProductsByUserID(user_id uint, term string, limit int, offset int) ([]Product, error)
	GetProductByID(id uint) (Product, users.User, error)
	UpdateProductByID(id uint, updatedProduct Product) error
	DeleteProduct(id uint) error
	CountAllProducts(term string) (int64, error)
	CountProductsByUserID(userID uint, term string) (int64, error)
}

type PQuery interface {
	AddProduct(newProduct Product) error
	GetAllProducts(term string, limit int, offset int) ([]Product, error)
	GetProductsByUserID(user_id uint, term string, limit int, offset int) ([]Product, error)
	GetProductByID(id uint) (Product, users.User, error)
	UpdateProductByID(id uint, updatedProduct Product) error
	DeleteProduct(id uint) error
	CountAllProducts(term string) (int64, error)
	CountProductsByUserID(userID uint, term string) (int64, error)
}
