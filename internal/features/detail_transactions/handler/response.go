package handler

import (
	dt_entity "TokoGadget/internal/features/detail_transactions"
	t_entity "TokoGadget/internal/features/transactions"
	"time"
)

type CartResponse struct {
	TransactionID uint                              `json:"transaction_id"`
	UserID        uint                              `json:"user_id"`
	Status        string                            `json:"status"`
	CartItems     []dt_entity.AllDetailTransactions `json:"cart_items"`
	GrandTotal    uint64                            `json:"grand_total"`
	CreatedAt     time.Time                         `json:"created_at"`
	UpdatedAt     time.Time                         `json:"updated_at"`
	DeletedAt     time.Time                         `json:"deleted_at"`
	Meta          []Meta                            `json:"meta"`
}

type Meta struct {
	TotalItems   uint `json:"totalItems"`
	ItemsPerPage uint `json:"itemsPerPage"`
	CurrentPage  uint `json:"currentPage"`
	TotalPages   uint `json:"totalPages"`
}

func ToCartResponse(t t_entity.Transaction, dt []dt_entity.AllDetailTransactions) CartResponse {
	// Hitung Grand Total
	var grandTotal uint64
	for _, val := range dt {
		grandTotal += val.SubTotal
	}

	return CartResponse{
		TransactionID: t.ID,
		UserID:        t.UserID,
		Status:        t.Status,
		CartItems:     dt,
		GrandTotal:    grandTotal,
		CreatedAt:     t.CreatedAt,
		UpdatedAt:     t.UpdatedAt,
		DeletedAt:     t.DeletedAt,
	}
}
