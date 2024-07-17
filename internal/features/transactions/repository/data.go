package repository

import (
	dt_rep "TokoGadget/internal/features/detail_transactions/repository"
	t_entity "TokoGadget/internal/features/transactions"
	u_rep "TokoGadget/internal/features/users/repository"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID             uint
	Status             string
	Users              u_rep.User               `gorm:"foreignKey:UserID"`
	DetailTransactions dt_rep.DetailTransaction `gorm:"foreignKey:TransactionID"`
}

func ToEntityTransaction(input Transaction) t_entity.Transaction {
	return t_entity.Transaction{
		ID:     input.ID,
		UserID: input.UserID,
		Status: input.Status,
	}
}
