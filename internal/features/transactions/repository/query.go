package repository

import (
	p_qry "TokoGadget/internal/features/products/repository"
	"TokoGadget/internal/features/transactions"
	"fmt"

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
	query := tq.db.Raw("SELECT dt.id AS cart_id, p.id AS product_id, dt.quantity, p.stock FROM detail_transactions AS dt JOIN products AS p ON p.id = dt.product_id WHERE transaction_id = ?", transactionID)
	query.Scan(&stock)

	// Stok mencukupi = true, tidak cukup = false
	for _, val := range stock {
		if val.Quantity > val.Stock {
			return []transactions.CheckStock{}, false
		}
	}

	return stock, true
}


func (tq *TransactionQuery) GetPaymentDetails(transactionID uint) transactions.PaymentDetails {
	// Get Fullname
	var result transactions.PaymentDetails
	query := tq.db.Raw("SELECT t.id AS transaction_id, u.fullname FROM transactions AS t JOIN users AS u ON t.user_id = u.id WHERE t.id = ?", transactionID)
	query.Scan(&result)
	fmt.Println("Datanya : ", result)

	// Get All Cart Data Subtotals
	var AllCart []CartSubTotals
	query2 := tq.db.Raw("SELECT p.price AS price, dt.quantity FROM detail_transactions AS dt JOIN products AS p ON dt.product_id = p.id WHERE transaction_id = ?", transactionID)
	query2.Scan(&AllCart)
	fmt.Println("Data Query :", AllCart)

	// Calculate Total Amount
	var amount uint64
	for _, val := range AllCart {
		amount += val.Price * val.Quantity
	}
	result.Ammount = amount
	fmt.Println("Data result sebelum return : ", result)
	return result
}


func (tq *TransactionQuery) UpdateStock(input []transactions.CheckStock) error {
	for _, val := range input {
		newStock := val.Stock - val.Quantity
		fmt.Println("Stock terbaru :", newStock)
		fmt.Println("id produk :", val.ProductID)
		err := tq.db.Model(&p_qry.Product{}).Where("id = ?", val.ProductID).UpdateColumn("stock", newStock).Error
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
	var trans []Transaction
	err := tq.db.Where("user_id = ? AND status = 'success'", userID).Find(&trans).Error
	if err != nil {
		return []transactions.Transaction{}, err
	}

	transCnv := ToAllEntityTransaction(trans)

	for i, val := range trans {
		var result []transactions.AllDetailTransactions
		query := tq.db.Raw("SELECT dt.id as cart_id, p.product_name, p.product_picture, dt.quantity, p.price AS sub_total FROM detail_transactions AS dt JOIN transactions AS t ON t.id = dt.transaction_id JOIN products AS p ON p.id = dt.product_id WHERE t.id = ? AND t.status = 'success' AND dt.deleted_at IS NULL", val.ID)
		err = query.Scan(&result).Error
		if err != nil {
			return []transactions.Transaction{}, err
		}

		for i, v := range result {
			result[i].SubTotal = uint64(v.Quantity) * uint64(v.SubTotal)
		}

		transCnv[i].TransactionItems = append(transCnv[i].TransactionItems, result...)

		for _, v := range result {
			transCnv[i].GrandTotal += v.SubTotal
		}
	}

	return transCnv, nil
}

func (tq *TransactionQuery) GetTransaction(transactionID uint) (transactions.Transaction, error) {
	var result Transaction
	err := tq.db.Where("id = ?", transactionID).First(&result).Error
	if err != nil {
		return transactions.Transaction{}, err
	}

	resultCnv := ToEntityTransaction(result)

	var result2 []transactions.AllDetailTransactions
	query := tq.db.Raw("SELECT dt.id as cart_id, p.product_name, p.product_picture, dt.quantity, p.price AS sub_total FROM detail_transactions AS dt JOIN transactions AS t ON t.id = dt.transaction_id JOIN products AS p ON p.id = dt.product_id WHERE t.id = ? AND t.status = 'success' AND dt.deleted_at IS NULL", transactionID)
	err = query.Scan(&result2).Error
	if err != nil {
		return transactions.Transaction{}, err
	}

	resultCnv.TransactionItems = result2

	for i, v := range result2 {
		result2[i].SubTotal = uint64(v.Quantity) * uint64(v.SubTotal)
	}

	for _, v := range result2 {
		resultCnv.GrandTotal += v.SubTotal
	}

	return resultCnv, nil
}

func (tq *TransactionQuery) DeleteTransaction(transactionID uint) error {
	err := tq.db.Delete(&Transaction{}, transactionID).Error
	if err != nil {
		return err
	}

	return nil
}

func (tq *TransactionQuery) RevertStock(input []transactions.CheckStock) error {
	for _, val := range input {
		newStock := val.Stock + val.Quantity
		err := tq.db.Model(&p_qry.Product{}).Where("id = ?", val.ProductID).UpdateColumn("stock", newStock).Error
		if err != nil {
			return err
		}
	}

	return nil
}
