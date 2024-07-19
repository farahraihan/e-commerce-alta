package services_test

import (
	dt_entity "TokoGadget/internal/features/detail_transactions"
	"TokoGadget/internal/features/detail_transactions/services"
	t_entity "TokoGadget/internal/features/transactions"
	"TokoGadget/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func ImportMocks(t *testing.T) (*mocks.DTQuery, *mocks.TQuery, dt_entity.DTServices) {
	dt_qry := mocks.NewDTQuery(t)
	t_qry := mocks.NewTQuery(t)
	srv := services.NewDetailTransactionServices(dt_qry, t_qry)

	return dt_qry, t_qry, srv
}

func TestGetAllCart(t *testing.T) {
	dt_qry, t_qry, srv := ImportMocks(t)

	var userID uint = 1

	t.Run("CreateTransaction Error", func(t *testing.T) {
		var check t_entity.Transaction

		t_qry.On("CheckPendingTransaction", userID).Return(check, nil).Once()
		t_qry.On("CreateTransaction", userID).Return(gorm.ErrInvalidData).Once()

		_, _, err := srv.GetAllCart(userID)

		t_qry.AssertExpectations(t)
		assert.Error(t, err)
	})

	t.Run("GetAllCart Success", func(t *testing.T) {
		var check t_entity.Transaction
		var cart []dt_entity.AllDetailTransactions

		t_qry.On("CheckPendingTransaction", userID).Return(check, nil).Once()
		t_qry.On("CreateTransaction", userID).Return(nil).Once()
		t_qry.On("CheckPendingTransaction", userID).Return(check, nil).Once()
		dt_qry.On("GetAllCart", userID).Return(cart, nil).Once()

		_, _, err := srv.GetAllCart(userID)

		t_qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestAddToCart(t *testing.T) {
	dt_qry, t_qry, srv := ImportMocks(t)

	var productID uint = 1
	var userID uint = 1

	t.Run("AddToCart Success", func(t *testing.T) {
		var result t_entity.Transaction
		t_qry.On("CheckPendingTransaction", userID).Return(result, nil).Once()
		t_qry.On("CreateTransaction", userID).Return(nil).Once()
		dt_qry.On("AddToCart", productID, userID).Return(nil).Once()

		err := srv.AddToCart(productID, userID)

		t_qry.AssertExpectations(t)
		dt_qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestUpdateCart(t *testing.T) {
	dt_qry, _, srv := ImportMocks(t)

	var cartID uint = 1
	var quantity uint = 1

	t.Run("UpdateCart Success", func(t *testing.T) {
		dt_qry.On("CheckStockPerProduct", cartID, quantity).Return(true).Once()
		dt_qry.On("UpdateCart", cartID, quantity).Return(nil).Once()

		_, err := srv.UpdateCart(cartID, quantity)

		dt_qry.AssertExpectations(t)
		assert.Nil(t, err)
	})
}

func TestDeleteCart(t *testing.T) {
	dt_qry, _, srv := ImportMocks(t)

	var cartID uint = 1

	t.Run("DeleteCart Success", func(t *testing.T) {
		dt_qry.On("DeleteCart", cartID).Return(nil).Once()

		err := srv.DeleteCart(cartID)

		dt_qry.AssertExpectations(t)
		assert.Nil(t, err)
	})

}
