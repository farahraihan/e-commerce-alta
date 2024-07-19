package repository

import (
	detail_transaction "TokoGadget/internal/features/detail_transactions"
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/transactions"
	"TokoGadget/internal/features/users"
)

type Sales struct {
	UserID              uint                                 `json:"seller_id"`
	User                users.User                           `gorm:"foreignKey:SellerID"`
	ProductID           uint                                 `json:"product_id"`
	Product             products.Product                     `gorm:"foreignKey:ProductID"`
	TransactionID       uint                                 `json:"transaction_id"`
	Transaction         transactions.Transaction             `gorm:"foreignKey:TransactionID"`
	DetailTransactionID uint                                 `json:"detail_transaction_id"`
	DetailTransaction   detail_transaction.DetailTransaction `gorm:"foreignKey:DetailTransactionID"`
}
