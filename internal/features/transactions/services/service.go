package services

import (
	t_entity "TokoGadget/internal/features/transactions"
)

type TransactionServices struct {
	qry t_entity.TQuery
}

func NewTransactionServices(q t_entity.TQuery) t_entity.TServices {
	return &TransactionServices{
		qry: q,
	}
}

func (ts *TransactionServices) Checkout(transactionID uint) error {
	return ts.qry.Checkout(transactionID)
}

func (ts *TransactionServices) GetAllTransactions(userID uint) ([]t_entity.Transaction, error) {
	return ts.qry.GetAllTransactions(userID)
}

func (ts *TransactionServices) GetTransaction(transactionID uint) (t_entity.Transaction, error) {
	return ts.qry.GetTransaction(transactionID)
}

func (ts *TransactionServices) DeleteTransaction(transactionID uint) error {
	return ts.qry.DeleteTransaction(transactionID)
}
