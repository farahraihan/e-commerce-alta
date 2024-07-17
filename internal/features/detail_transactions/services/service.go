package services

import (
	dt_entity "TokoGadget/internal/features/detail_transactions"
	t_entity "TokoGadget/internal/features/transactions"
	// t_rep "TokoGadget/internal/features/transactions/repository"
)

type DetailTransactionServices struct {
	qry       dt_entity.DTQuery
	qry_trans t_entity.TQuery
}

func NewDetailTransactionServices(q dt_entity.DTQuery, qt t_entity.TQuery) dt_entity.DTServices {
	return &DetailTransactionServices{
		qry:       q,
		qry_trans: qt,
	}
}

func (dts *DetailTransactionServices) GetAllCart(userID uint) (t_entity.Transaction, []dt_entity.AllDetailTransactions, error) {
	check, _ := dts.qry_trans.CheckPendingTransaction(userID)
	if check.ID == 0 {
		err := dts.qry_trans.CreateTransaction(userID)
		if err != nil {
			return t_entity.Transaction{}, []dt_entity.AllDetailTransactions{}, err
		}
	}

	transaction, _ := dts.qry_trans.CheckPendingTransaction(userID)
	cart, err := dts.qry.GetAllCart(userID)

	// Hitung SubTotal
	for i, val := range cart {
		subTotal := val.SubTotal * uint64(val.Quantity)
		cart[i].SubTotal = subTotal
	}

	return transaction, cart, err
}

func (dts *DetailTransactionServices) AddToCart(productID uint, userID uint) error {
	result, err := dts.qry_trans.CheckPendingTransaction(userID)
	if err != nil {
		return err
	}

	if result.ID == 0 {
		err := dts.qry_trans.CreateTransaction(userID)
		if err != nil {
			return err
		}
	}

	return dts.qry.AddToCart(productID, userID)
}

func (dts *DetailTransactionServices) UpdateCart(cartID uint, quantity uint) (bool, error) {
	status := dts.qry.CheckStockPerProduct(cartID, quantity)
	if !status {
		return false, nil
	}

	return true, dts.qry.UpdateCart(cartID, quantity)
}

func (dts *DetailTransactionServices) DeleteCart(cartID uint) error {
	return dts.qry.DeleteCart(cartID)
}
