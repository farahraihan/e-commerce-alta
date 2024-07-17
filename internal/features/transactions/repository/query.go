package repository

import (
	p_qry "TokoGadget/internal/features/products/repository"
	"TokoGadget/internal/features/transactions"

	"gorm.io/gorm"
)

type TransactionQuery struct {
	db *gorm.DB
}

func NewTransactionQuery(connection *gorm.DB) transactions.TQuery {
	return &TransactionQuery{
		db: connection,
	}
}

func (tq *TransactionQuery) CreateTransaction(userID uint) error {
	input := Transaction{
		UserID: userID,
		Status: "pending",
	}

	err := tq.db.Create(&input).Error
	if err != nil {
		return err
	}

	return nil
}

func (tq *TransactionQuery) CheckPendingTransaction(userID uint) (transactions.Transaction, error) {
	var result Transaction
	err := tq.db.Where("user_id = ? AND status = 'pending'", userID).First(&result).Error
	if err != nil {
		return transactions.Transaction{}, err
	}

	return ToEntityTransaction(result), nil
}

func (tq *TransactionQuery) CheckStock(transactionID uint) ([]transactions.CheckStock, bool) {
	var stock []transactions.CheckStock
	query := tq.db.Raw("SELECT dt.id AS cart_id, dt.quantity, p.stock FROM detail_transactions AS dt JOIN products AS p ON p.id = dt.product_id WHERE transaction_id = ?", transactionID)
	query.Scan(&stock)

	// Stok mencukupi = true, tidak cukup = false
	for _, val := range stock {
		if val.Quantity > val.Stock {
			return []transactions.CheckStock{}, false
		}
	}

	return stock, true
}

func (tq *TransactionQuery) UpdateStock(input []transactions.CheckStock) error {
	for _, val := range input {
		newStock := val.Stock - val.Quantity
		err := tq.db.Model(&p_qry.Product{}).Where("product_id = ?", val.ProductID).UpdateColumn("stock", newStock).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (tq *TransactionQuery) Checkout(transactionID uint) error {
	err := tq.db.Model(&Transaction{}).Where("id = ?", transactionID).UpdateColumn("status", "success").Error
	if err != nil {
		return err
	}

	return nil
}

func (tq *TransactionQuery) GetAllTransactions(userID uint) ([]transactions.Transaction, error) {
	var result []transactions.Transaction
	err := tq.db.Where("user_id = ? AND status = 'success'", userID).Find(&result).Error
	if err != nil {
		return []transactions.Transaction{}, err
	}

	return result, nil
}

func (tq *TransactionQuery) GetTransaction(transactionID uint) (transactions.Transaction, error) {
	var result transactions.Transaction
	err := tq.db.Where("transaction_id = ?", transactionID).First(&result).Error
	if err != nil {
		return transactions.Transaction{}, err
	}

	return result, nil
}

func (tq *TransactionQuery) DeleteTransaction(transactionID uint) error {
	err := tq.db.Delete(&transactions.Transaction{}, transactionID).Error
	if err != nil {
		return err
	}

	return nil
}

func (tq *TransactionQuery) RevertStock(input []transactions.CheckStock) error {
	for _, val := range input {
		newStock := val.Stock + val.Quantity
		err := tq.db.Model(&p_qry.Product{}).Where("product_id = ?", val.ProductID).UpdateColumn("stock", newStock).Error
		if err != nil {
			return err
		}
	}

	return nil
}
