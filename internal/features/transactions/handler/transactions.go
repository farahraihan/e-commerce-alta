package handler

import (
	t_entity "TokoGadget/internal/features/transactions"
	"TokoGadget/internal/helper"
	"TokoGadget/internal/utils"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	srv t_entity.TServices
	tu  utils.TokenUtilityInterface
}

func NewTransactionHandler(s t_entity.TServices, t utils.TokenUtilityInterface) t_entity.THandler {
	return &TransactionHandler{
		srv: s,
		tu:  t,
	}
}
func (th *TransactionHandler) Checkout(c echo.Context) error {
    // Parse transaction_id from URL parameter
    orderID, err := strconv.ParseUint(c.Param("transaction_id"), 10, 64)
    if err != nil {
        return c.JSON(400, helper.ResponseFormat("Failed", 400, "Input Error!", nil, nil))
    }

    // Request Midtrans payment URL
    redirectURL, err := th.srv.RequestMidtransPayment(uint(orderID))
    if err != nil {
        return c.JSON(500, helper.ResponseFormat("Failed", 500, err.Error(), nil, nil))
    }

    if redirectURL == "" {
        return c.JSON(500, helper.ResponseFormat("Failed", 500, "Empty redirect URL from Midtrans!", nil, nil))
    }

    // Prepare PaymentResponse struct
    response := PaymentResponse{
        RedirectURL: redirectURL,
    }

    // Return success response with PaymentResponse
    return c.JSON(200, helper.ResponseFormat("Success", 200, "Checkout successful!", response, nil))
}

func (th *TransactionHandler) CheckStatusPayment(c echo.Context) error {
	// Bind request body to TransactionStatusRequest struct
	var req TransactionStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, helper.ResponseFormat("Failed", 400, "Invalid request payload", nil, nil))
	}
	fmt.Println("Order ID :", req.OrderID)
	orderID, err := strconv.ParseUint(req.OrderID, 10, 64)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat("Failed", 400, "Invalid Order ID", nil, nil))
	}

	if req.TransactionStatus != "settlement" {
		return c.JSON(400, helper.ResponseFormatNonData(400, req.TransactionStatus, "Pembayaran belum selesai"))
	}
	// Checkout transaction after payment request
	_, paymentStatus, err := th.srv.Checkout(uint(orderID))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat("Failed", 500, err.Error(), nil, nil))
	}

	// // Check payment status
	if !paymentStatus {
		return c.JSON(400, helper.ResponseFormat("Failed", 400, "Payment failed!", nil, nil))
	}

	// Payment successful response
	response := map[string]interface{}{
		"transaction_time":   req.TransactionTime,
		"transaction_status": req.TransactionStatus,
		"transaction_id":     req.TransactionID,
		"order_id":           req.OrderID,
	}
	return c.JSON(200, helper.ResponseFormat("success", 200, "Pembayaran Berhasil", response, nil))
}



func (th *TransactionHandler) GetAllTransactions(c echo.Context) error {
	userID := th.tu.DecodeToken(c.Get("user").(*jwt.Token))
	result, err := th.srv.GetAllTransactions(userID)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat("Failed", 500, "Server Error!", nil, nil))
	}

	return c.JSON(200, helper.ResponseFormat("Failed", 200, "All Transactions History successfully retreived!", result, nil))
}

func (th *TransactionHandler) GetTransaction(c echo.Context) error {
	transactionID, _ := strconv.Atoi(c.Param("transaction_id"))
	result, err := th.srv.GetTransaction(uint(transactionID))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat("Failed", 500, "Server Error!", nil, nil))
	}

	return c.JSON(200, helper.ResponseFormat("Failed", 200, "Transaction History successfully retreived!", result, nil))
}

func (th *TransactionHandler) DeleteTransaction(c echo.Context) error {
	transactionID, _ := strconv.Atoi(c.Param("transaction_id"))
	err := th.srv.DeleteTransaction(uint(transactionID))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat("Failed", 500, "Server Error!", nil, nil))
	}

	return c.JSON(200, helper.ResponseFormat("Failed", 200, "Transaction successfully canceled!", nil, nil))
}
