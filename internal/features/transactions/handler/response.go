package handler

import (
	"TokoGadget/internal/features/transactions"
	"time"
)

type TransactionResponse struct {
	TransactionID    uint                                 `json:"transaction_id"`
	UserID           uint                                 `json:"user_id"`
	Status           string                               `json:"status"`
	TransactionItems []transactions.AllDetailTransactions `json:"transaction_items"`
	GrandTotal       uint64                               `json:"grand_total"`
	CreatedAt        time.Time                            `json:"created_at"`
	UpdatedAt        time.Time                            `json:"updated_at"`
	DeletedAt        time.Time                            `json:"deleted_at"`
}
type PaymentResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

type TransactionStatusRequest struct {
	TransactionTime   string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	TransactionID     string `json:"transaction_id"`
	OrderID           string `json:"order_id"`
}
