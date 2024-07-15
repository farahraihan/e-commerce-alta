package repository

import (
	detailtransactions "TokoGadget/internal/features/detail_transactions"
	"TokoGadget/internal/features/users"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID             uint                                   `json:"user_id"`
	User               users.User                             `gorm:"foreignKey:UserID"`
	Status             string                                 `json:"status"`
	detailtransactions []detailtransactions.DetailTransaction `gorm:"foreignKey:TransactionID"`
}
