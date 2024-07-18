package services

import (
	t_entity "TokoGadget/internal/features/transactions"
	"TokoGadget/internal/utils"
)

type TransactionServices struct {
	qry t_entity.TQuery
	mi  utils.MidtransInterface
}

func NewTransactionServices(q t_entity.TQuery, m utils.MidtransInterface) t_entity.TServices {
	return &TransactionServices{
		qry: q,
		mi:  m,
	}
}

func (ts *TransactionServices) Checkout(transactionID uint) (bool, error) {
	// Check stock
	result, status := ts.qry.CheckStock(transactionID)
	if !status {
		return false, nil
	}

	// // Get Transaction Details
	// paymentDetails := ts.qry.GetPaymentDetails(transactionID)

	// // Payment Gateway
	// _, err := ts.mi.RequestPayment(strconv.Itoa(int(transactionID)), int64(paymentDetails.Ammount))
	// if err != nil {
	// 	return true, false, err
	// }

	// Update Product Stock After Payment Success
	err := ts.qry.UpdateStock(result)
	if err != nil {
		return true, err
	}

	// Update Transaction Status to True
	return true, ts.qry.Checkout(transactionID)
}

func (ts *TransactionServices) GetAllTransactions(userID uint) ([]t_entity.Transaction, error) {
	return ts.qry.GetAllTransactions(userID)
}

func (ts *TransactionServices) GetTransaction(transactionID uint) (t_entity.Transaction, error) {
	return ts.qry.GetTransaction(transactionID)
}

func (ts *TransactionServices) DeleteTransaction(transactionID uint) error {
	result, _ := ts.qry.CheckStock(transactionID)
	err := ts.qry.RevertStock(result)
	if err != nil {
		return err
	}

	return ts.qry.DeleteTransaction(transactionID)
}
