package sales

import (
	detail_transaction "TokoGadget/internal/features/detail_transactions"
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/transactions"
	"TokoGadget/internal/features/users"

	"github.com/labstack/echo/v4"
)

type Sale struct {
	UserID              uint
	ProductID           uint
	TransactionID       uint
	DetailTransactionID uint
}

type SHandler interface {
	GetSalesByUserID() echo.HandlerFunc
	GetSalesByTransactionID() echo.HandlerFunc
}

type SServices interface {
	GetSalesByUserID(UserID uint) (users.User, []products.Product, []transactions.Transaction, []detail_transaction.DetailTransaction, error)
	GetSalesByTransactionID(UserID uint, TransactionID uint) (users.User, []products.Product, transactions.Transaction, []detail_transaction.DetailTransaction, error)
}

type SQuery interface {
	GetSalesByUserID(UserID uint) (users.User, []products.Product, []transactions.Transaction, []detail_transaction.DetailTransaction, error)
	GetSalesByTransactionID(UserID uint, TransactionID uint) (users.User, []products.Product, transactions.Transaction, []detail_transaction.DetailTransaction, error)
}
