package services_test

import (
	t_entity "TokoGadget/internal/features/transactions"
	"TokoGadget/internal/features/transactions/services"
	"TokoGadget/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ImportMocks(t *testing.T) (*mocks.TQuery, *mocks.MidtransInterface, t_entity.TServices) {
	qry := mocks.NewTQuery(t)
	mi := mocks.NewMidtransInterface(t)
	srv := services.NewTransactionServices(qry, mi)

	return qry, mi, srv
}

func TestCheckout(t *testing.T) {
	qry, _, srv := ImportMocks(t)

	var transactionID uint = 1

	t.Run("Checkout Success", func(t *testing.T) {
		var result []t_entity.CheckStock

		qry.On("CheckStock", transactionID).Return(result, true).Once()
		qry.On("UpdateStock", result).Return(nil).Once()
		qry.On("Checkout", transactionID).Return(nil).Once()

		_, _, err := srv.Checkout(transactionID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestRequestMidtransPayment(t *testing.T) {
	qry, mi, srv := ImportMocks(t)

	var transactionID uint = 1

	t.Run("RequestMidtransPayment Success", func(t *testing.T) {
		var paymentDetails t_entity.PaymentDetails

		qry.On("GetPaymentDetails", transactionID).Return(paymentDetails).Once()
		mi.On("RequestPayment", "1", 0).Return("link", nil).Once()

		_, err := srv.RequestMidtransPayment(transactionID)

		qry.AssertExpectations(t)
		mi.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestGetAllTransactions(t *testing.T) {
	qry, _, srv := ImportMocks(t)

	var userID uint = 1

	t.Run("GetAllTransactions Success", func(t *testing.T) {
		var result []t_entity.Transaction

		qry.On("GetAllTransactions", userID).Return(result, nil).Once()

		_, err := srv.GetAllTransactions(userID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestGetTransaction(t *testing.T) {
	qry, _, srv := ImportMocks(t)

	var transactionID uint = 1

	t.Run("GetTransaction Success", func(t *testing.T) {
		var result t_entity.Transaction

		qry.On("GetTransaction", transactionID).Return(result, nil).Once()

		_, err := srv.GetTransaction(transactionID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestDeleteTransaction(t *testing.T) {
	qry, _, srv := ImportMocks(t)

	var transactionID uint = 1

	t.Run("DeleteTransaction Success", func(t *testing.T) {
		var result []t_entity.CheckStock

		qry.On("CheckStock", transactionID).Return(result, true).Once()
		qry.On("RevertStock", result).Return(nil).Once()
		qry.On("DeleteTransaction", transactionID).Return(nil).Once()

		err := srv.DeleteTransaction(transactionID)

		qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}
