package repository

import (
	detail_transaction "TokoGadget/internal/features/detail_transactions"
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/sales"
	"TokoGadget/internal/features/transactions"
	"TokoGadget/internal/features/users"

	"gorm.io/gorm"
)

type SaleModel struct {
	db *gorm.DB
}

func NewSaleModel(connection *gorm.DB) sales.SQuery {
	return &SaleModel{
		db: connection,
	}
}

func (s *SaleModel) GetSalesByUserID(UserID uint) (users.User, []products.Product, []transactions.Transaction, []detail_transaction.DetailTransaction, error) {
	var user users.User
	var products []products.Product
	var transactions []transactions.Transaction
	var detailTransactions []detail_transaction.DetailTransaction

	// Ambil penjual berdasarkan ID
	if err := s.db.First(&user, UserID).Error; err != nil {
		return users.User{}, nil, nil, nil, err
	}

	// Ambil semua produk yang dimiliki atau dijual oleh penjual
	if err := s.db.Where("user_id = ?", UserID).Find(&products).Error; err != nil {
		return users.User{}, nil, nil, nil, err
	}

	// Ambil semua detail transaksi yang terkait dengan produk-produk tersebut
	productIDs := make([]uint, len(products))
	for i, product := range products {
		productIDs[i] = product.ID
	}
	if err := s.db.Where("product_id IN (?)", productIDs).Find(&detailTransactions).Error; err != nil {
		return users.User{}, nil, nil, nil, err
	}

	// Ambil semua transaksi yang terkait dengan detail transaksi tersebut
	transactionIDs := make([]uint, len(detailTransactions))
	for i, detailTransaction := range detailTransactions {
		transactionIDs[i] = detailTransaction.TransactionID
	}
	if err := s.db.Where("id IN (?)", transactionIDs).Find(&transactions).Error; err != nil {
		return users.User{}, nil, nil, nil, err
	}

	return user, products, transactions, detailTransactions, nil
}

func (s *SaleModel) GetSalesByTransactionID(UserID uint, TransactionID uint) (users.User, []products.Product, transactions.Transaction, []detail_transaction.DetailTransaction, error) {
	var user users.User
	var products []products.Product
	var transaction transactions.Transaction
	var detailTransactions []detail_transaction.DetailTransaction

	// Ambil pengguna (penjual) berdasarkan ID
	if err := s.db.First(&user, UserID).Error; err != nil {
		return users.User{}, nil, transactions.Transaction{}, nil, err
	}

	// Ambil semua produk yang dimiliki oleh penjual
	if err := s.db.Where("user_id = ?", UserID).Find(&products).Error; err != nil {
		return users.User{}, nil, transactions.Transaction{}, nil, err
	}

	// Ambil semua detail transaksi berdasarkan produk-produk tersebut dan TransactionID
	productIDs := make([]uint, len(products))
	for i, product := range products {
		productIDs[i] = product.ID
	}
	if err := s.db.Where("product_id IN (?) AND transaction_id = ?", productIDs, TransactionID).Find(&detailTransactions).Error; err != nil {
		return users.User{}, nil, transactions.Transaction{}, nil, err
	}

	// Pastikan setidaknya ada satu detail transaksi yang ditemukan
	if len(detailTransactions) == 0 {
		return users.User{}, nil, transactions.Transaction{}, nil, gorm.ErrRecordNotFound
	}

	// Ambil transaksi berdasarkan TransactionID
	if err := s.db.First(&transaction, TransactionID).Error; err != nil {
		return users.User{}, nil, transactions.Transaction{}, nil, err
	}

	return user, products, transaction, detailTransactions, nil
}
