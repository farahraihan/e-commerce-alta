package handler

import (
	t_entity "TokoGadget/internal/features/transactions"
	"TokoGadget/internal/helper"
	"TokoGadget/internal/utils"
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
	var orderID uint
	err := c.Bind(&orderID)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "Input Error!", nil, nil))
	}

	stockStatus, paymentStatus, err := th.srv.Checkout(orderID)
	if !stockStatus {
		return c.JSON(400, helper.ResponseFormat(400, "Not enough stock!", nil, nil))
	}
	if !paymentStatus {
		return c.JSON(400, helper.ResponseFormat(400, "Payment failed!", nil, nil))
	}
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "Server Error!", nil, nil))
	}

	return c.JSON(200, helper.ResponseFormat(201, "All item in the Shopping Cart has been successfully checked out!", nil, nil))
}

func (th *TransactionHandler) GetAllTransactions(c echo.Context) error {
	userID := th.tu.DecodeToken(c.Get("user").(*jwt.Token))
	result, err := th.srv.GetAllTransactions(userID)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "Server Error!", nil, nil))
	}

	return c.JSON(200, helper.ResponseFormat(200, "All Transactions History successfully retreived!", result, nil))
}

func (th *TransactionHandler) GetTransaction(c echo.Context) error {
	transactionID, _ := strconv.Atoi(c.Param("transaction_id"))
	result, err := th.srv.GetTransaction(uint(transactionID))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "Server Error!", nil, nil))
	}

	return c.JSON(200, helper.ResponseFormat(200, "Transaction History successfully retreived!", result, nil))
}

func (th *TransactionHandler) DeleteTransaction(c echo.Context) error {
	transactionID, _ := strconv.Atoi(c.Param("transaction_id"))
	err := th.srv.DeleteTransaction(uint(transactionID))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "Server Error!", nil, nil))
	}

	return c.JSON(200, helper.ResponseFormat(200, "Transaction successfully canceled!", nil, nil))
}
