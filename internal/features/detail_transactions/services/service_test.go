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
	_, t_qry, srv := ImportMocks(t)

	var userID uint = 1
	// var errExample error

	t.Run("Get All Cart Success", func(t *testing.T) {
		var check t_entity.Transaction
		t_qry.On("CheckPendingTransaction", userID).Return(check).Once()
		t_qry.On("CreateTransaction", userID).Return(gorm.ErrInvalidData).Once()

		_, _, err := srv.GetAllCart(userID)

		t_qry.AssertExpectations(t)
		assert.Error(t, err)
	})
}
