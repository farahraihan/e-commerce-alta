package transactions

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Transaction struct {
	ID        uint
	UserID    uint
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type CheckStock struct {
	CartID    uint
	ProductID uint
	Quantity  uint
	Stock     uint
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
	UpdateStock([]CheckStock) error
	Checkout(uint) error
	GetAllTransactions(uint) ([]Transaction, error)
	GetTransaction(uint) (Transaction, error)
	DeleteTransaction(uint) error
	RevertStock([]CheckStock) error
}
