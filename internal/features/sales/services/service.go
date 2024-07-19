package services

import (
	detail_transaction "TokoGadget/internal/features/detail_transactions"
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/sales"
	"TokoGadget/internal/features/transactions"
	"TokoGadget/internal/features/users"
)

type SaleServices struct {
	qry sales.SQuery
}

func NewSaleService(q sales.SQuery) sales.SServices {
	return &SaleServices{
		qry: q,
	}
}

func (s *SaleServices) GetSalesByUserID(UserID uint) (users.User, []products.Product, []transactions.Transaction, []detail_transaction.DetailTransaction, error) {
	user, products, transactions, detailTransactions, err := s.qry.GetSalesByUserID(UserID)
	if err != nil {
		return users.User{}, nil, nil, nil, err
	}
	return user, products, transactions, detailTransactions, nil
}

func (s *SaleServices) GetSalesByTransactionID(UserID uint, TransactionID uint) (users.User, []products.Product, transactions.Transaction, []detail_transaction.DetailTransaction, error) {
	user, products, transaction, detailTransactions, err := s.qry.GetSalesByTransactionID(UserID, TransactionID)
	if err != nil {
		return users.User{}, nil, transactions.Transaction{}, nil, err
	}
	return user, products, transaction, detailTransactions, nil
}
