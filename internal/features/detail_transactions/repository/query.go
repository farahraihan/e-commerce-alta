package repository

import (
	dt_entity "TokoGadget/internal/features/detail_transactions"
	t_entity "TokoGadget/internal/features/transactions"

	"gorm.io/gorm"
)

type DetailTransactionQuery struct {
	db *gorm.DB
}

func NewDetailTransactionQuery(connection *gorm.DB) dt_entity.DTQuery {
	return &DetailTransactionQuery{
		db: connection,
	}
}

func (dtq *DetailTransactionQuery) GetAllCart(userID uint) ([]dt_entity.AllDetailTransactions, error) {
	var result []dt_entity.AllDetailTransactions

	query := dtq.db.Raw("SELECT dt.id as cart_id, p.product_name, p.product_picture, dt.quantity, p.price AS sub_total FROM detail_transactions AS dt JOIN transactions AS t ON t.id = dt.transaction_id JOIN products AS p ON p.id = dt.product_id WHERE t.user_id = ? AND t.status = 'pending' AND dt.deleted_at IS NULL", userID)
	err := query.Scan(&result).Error
	if err != nil {
		return []dt_entity.AllDetailTransactions{}, err
	}

	return result, nil
}

func (dtq *DetailTransactionQuery) AddToCart(productID uint, userID uint) error {
	var transactionID uint

	query := dtq.db.Raw("SELECT id FROM transactions WHERE user_id = ? AND status='pending'", userID)
	err := query.Scan(&transactionID).Error
	if err != nil {
		return err
	}

	input := DetailTransaction{
		ProductID:     productID,
		TransactionID: transactionID,
		Quantity:      1,
	}
	err = dtq.db.Create(&input).Error
	if err != nil {
		return err
	}

	return nil
}

func (dtq *DetailTransactionQuery) UpdateCart(cartID uint, quantity uint) error {
	err := dtq.db.Model(&DetailTransaction{}).Where("id = ?", cartID).UpdateColumn("quantity", quantity).Error
	if err != nil {
		return err
	}

	return nil
}

func (dtq *DetailTransactionQuery) CheckStockPerProduct(cartID uint, quantity uint) bool {
	var stock t_entity.CheckStock

	query := dtq.db.Raw("SELECT dt.id AS cart_id, dt.quantity, p.stock FROM detail_transactions AS dt JOIN products AS p ON p.id = dt.product_id WHERE dt.id = ?", cartID)
	query.Scan(&stock)

	return quantity <= stock.Stock
}

func (dtq *DetailTransactionQuery) DeleteCart(cartID uint) error {
	err := dtq.db.Delete(&DetailTransaction{}, cartID).Error
	if err != nil {
		return err
	}

	return nil
}
