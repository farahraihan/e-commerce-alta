package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	ID               uint                    `json:"transaction_id"`
	UserID           uint                    `json:"user_id"`
	Status           string                  `json:"status"`
	TransactionItems []AllDetailTransactions `json:"transaction_items"`
	GrandTotal       uint64                  `json:"grand_total"`
	CreatedAt        time.Time               `json:"created_at"`
	UpdatedAt        time.Time               `json:"updated_at"`
	DeletedAt        time.Time               `json:"deleted_at"`
}

type AllDetailTransactions struct {
	CartID         uint   `json:"cart_id"`
	ProductName    string `json:"product_name"`
	ProductPicture string `json:"product_picture"`
	Quantity       int    `json:"quantity"`
	SubTotal       uint64 `json:"sub_total"`
}

type CheckStock struct {
	CartID    uint
	ProductID uint
	Quantity  uint
	Stock     uint
}

type PaymentDetails struct {
	TransactionID uint
	Fullname      uint
	Ammount       uint64
}

type THandler interface {
	Checkout(echo.Context) error
	GetAllTransactions(echo.Context) error
	GetTransaction(echo.Context) error
	DeleteTransaction(echo.Context) error
}

type TServices interface {
	Checkout(uint) (bool, error)
	GetAllTransactions(uint) ([]Transaction, error)
	GetTransaction(uint) (Transaction, error)
	DeleteTransaction(uint) error
}

type TQuery interface {
	CreateTransaction(uint) error
	CheckPendingTransaction(uint) (Transaction, error)
	CheckStock(uint) ([]CheckStock, bool)
	GetPaymentDetails(uint) PaymentDetails
	UpdateStock([]CheckStock) error
	Checkout(uint) error
	GetAllTransactions(uint) ([]Transaction, error)
	GetTransaction(uint) (Transaction, error)
	DeleteTransaction(uint) error
	RevertStock([]CheckStock) error
}
