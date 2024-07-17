package services

import (
	dt_entity "TokoGadget/internal/features/detail_transactions"
	t_entity "TokoGadget/internal/features/transactions"

	// t_rep "TokoGadget/internal/features/transactions/repository"
	"strconv"
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

func (dts *DetailTransactionServices) GetAllCart(userID uint) ([]dt_entity.AllDetailTransactions, error) {
	result, err := dts.qry.GetAllCart(userID)

	// Hitung SubTotal
	for _, val := range result {
		price, _ := strconv.ParseInt(val.SubTotal, 10, 64)
		price *= int64(val.Quantity)
		val.SubTotal = strconv.FormatInt(price, 10)
	}

	return result, err
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

func (dts *DetailTransactionServices) UpdateCart(cartID uint, quantity uint) error {
	return dts.qry.UpdateCart(cartID, quantity)
}

func (dts *DetailTransactionServices) DeleteCart(cartID uint) error {
	return dts.qry.DeleteCart(cartID)
}
