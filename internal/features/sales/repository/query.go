package repository

import (
	detail_transaction "TokoGadget/internal/features/detail_transactions"
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/sales"
	"TokoGadget/internal/features/transactions"
	"TokoGadget/internal/features/users"

	"gorm.io/gorm"
)

type SalesModel struct {
	db *gorm.DB
}

func NewSalesModel(connection *gorm.DB) sales.SQuery {
	return &SalesModel{
		db: connection,
	}
}

func (s *SalesModel) GetSalesByUserID(UserID uint) (users.User, []products.Product, []transactions.Transaction, []detail_transaction.DetailTransaction, error) {
	var user users.User
	var products []products.Product
	var transactions []transactions.Transaction
	var detailTransactions []detail_transaction.DetailTransaction

	// Ambil penjual berdasarkan ID
	if err := s.db.First(&user, UserID).Error; err != nil {
		return users.User{}, nil, nil, nil, err
	}

	// Ambil semua produk yang dijual oleh penjual
	if err := s.db.Where("user_id = ?", UserID).Find(&products).Error; err != nil {
		return users.User{}, nil, nil, nil, err
	}

	// Ambil detail transaksi yang terkait dengan produk-produk tersebut
	productIDs := make([]uint, len(products))
	for i, product := range products {
		productIDs[i] = product.ID
	}
	if err := s.db.Where("product_id IN (?)", productIDs).Find(&detailTransactions).Error; err != nil {
		return users.User{}, nil, nil, nil, err
	}

	// Ambil transaksi yang terkait dengan detail transaksi tersebut
	transactionIDs := make([]uint, len(detailTransactions))
	for i, detailTransaction := range detailTransactions {
		transactionIDs[i] = detailTransaction.TransactionID
	}
	if err := s.db.Where("id IN (?)", transactionIDs).Find(&transactions).Error; err != nil {
		return users.User{}, nil, nil, nil, err
	}

	return user, products, transactions, detailTransactions, nil
}

func (s *SalesModel) GetSalesByTransactionID(UserID uint, TransactionID uint) (users.User, products.Product, transactions.Transaction, detail_transaction.DetailTransaction, error) {
	var user users.User
	var product products.Product
	var transaction transactions.Transaction
	var detailTransaction detail_transaction.DetailTransaction

	// Ambil produk berdasarkan TransactionID dan UserID
	if err := s.db.Table("products").
		Joins("JOIN detail_transactions ON products.id = detail_transactions.product_id").
		Joins("JOIN transactions ON detail_transactions.transaction_id = transactions.id").
		Where("products.user_id = ?", UserID).
		Where("transactions.id = ?", TransactionID).
		First(&product).Error; err != nil {
		return users.User{}, products.Product{}, transactions.Transaction{}, detail_transaction.DetailTransaction{}, err
	}

	// Ambil detail transaksi berdasarkan TransactionID
	if err := s.db.Where("transaction_id = ?", TransactionID).First(&detailTransaction).Error; err != nil {
		return users.User{}, products.Product{}, transactions.Transaction{}, detail_transaction.DetailTransaction{}, err
	}

	// Ambil transaksi berdasarkan TransactionID
	if err := s.db.First(&transaction, TransactionID).Error; err != nil {
		return users.User{}, products.Product{}, transactions.Transaction{}, detail_transaction.DetailTransaction{}, err
	}

	// Ambil pengguna (penjual) berdasarkan Product.UserID
	if err := s.db.First(&user, product.UserID).Error; err != nil {
		return users.User{}, products.Product{}, transactions.Transaction{}, detail_transaction.DetailTransaction{}, err
	}

	return user, product, transaction, detailTransaction, nil
}
