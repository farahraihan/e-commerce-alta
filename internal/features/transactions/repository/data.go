package repository

import (
	dt_rep "TokoGadget/internal/features/detail_transactions/repository"
	"TokoGadget/internal/features/sales"
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
	sales              []sales.Sales            `gorm:"foreignKey:TransactionID"`
}

func ToEntityTransaction(input Transaction) t_entity.Transaction {
	return t_entity.Transaction{
		ID:        input.ID,
		UserID:    input.UserID,
		Status:    input.Status,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		DeletedAt: input.DeletedAt.Time,
	}
}

func ToAllEntityTransaction(input []Transaction) []t_entity.Transaction {
	var result []t_entity.Transaction

	for _, val := range input {
		result = append(result, t_entity.Transaction{
			ID:        val.ID,
			UserID:    val.UserID,
			Status:    val.Status,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
			DeletedAt: val.DeletedAt.Time,
		})
	}

	return result
}

type CartSubTotals struct {
	Price    uint64
	Quantity uint64
}
