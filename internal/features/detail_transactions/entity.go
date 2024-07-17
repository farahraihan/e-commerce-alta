package detailtransactions

import "github.com/labstack/echo/v4"

type DetailTransaction struct {
	ID            uint
	ProductID     uint
	TransactionID uint
	Quantity      uint
}

type AllDetailTransactions struct {
	CartID         uint   `json:"cart_id"`
	ProductName    string `json:"product_name"`
	ProductPicture string `json:"product_picture"`
	Quantity       int    `json:"quantity"`
	SubTotal       string `json:"sub_total"`
}

type DTHandler interface {
	GetAllCart(echo.Context) error
	AddToCart(echo.Context) error
	UpdateCart(echo.Context) error
	DeleteCart(echo.Context) error
}

type DTServices interface {
	GetAllCart(uint) ([]AllDetailTransactions, error)
	AddToCart(uint, uint) error
	UpdateCart(uint, uint) error
	DeleteCart(uint) error
}

type DTQuery interface {
	GetAllCart(uint) ([]AllDetailTransactions, error)
	AddToCart(uint, uint) error
	UpdateCart(uint, uint) error
	DeleteCart(uint) error
}
