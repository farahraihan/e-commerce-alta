package handler

import (
	dt_entity "TokoGadget/internal/features/detail_transactions"
	"TokoGadget/internal/helper"
	"TokoGadget/internal/utils"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type DetailTransactionHandler struct {
	srv dt_entity.DTServices
	tu  utils.TokenUtilityInterface
}

func NewDetailTransactionHandler(s dt_entity.DTServices, t utils.TokenUtilityInterface) dt_entity.DTHandler {
	return &DetailTransactionHandler{
		srv: s,
		tu:  t,
	}
}

func (dth *DetailTransactionHandler) GetAllCart(c echo.Context) error {
	// Route: GET /cart
	userID := dth.tu.DecodeToken(c.Get("user").(*jwt.Token))

	transaction, cart, err := dth.srv.GetAllCart(uint(userID))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat("Failed", 500, "Server Error!", nil, nil))
	}

	return c.JSON(200, helper.ResponseFormat("Succeed", 200, "Shopping Cart datas successfully retreived!", ToCartResponse(transaction, cart), nil))
}

func (dth *DetailTransactionHandler) AddToCart(c echo.Context) error {
	// Route: POST /cart
	userID := dth.tu.DecodeToken(c.Get("user").(*jwt.Token))

	var input RequestCart
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat("Failed", 400, "Input Error!", nil, nil))
	}

	err = dth.srv.AddToCart(input.ProductID, userID)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat("Failed", 500, "Server Error!", nil, nil))
	}

	return c.JSON(201, helper.ResponseFormat("Succeed", 201, "Product successfully added to Shopping Cart!", nil, nil))
}

func (dth *DetailTransactionHandler) UpdateCart(c echo.Context) error {
	// Route: PUT /cart
	var input RequestCart
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat("Failed", 400, "Input Error!", nil, nil))
	}

	status, err := dth.srv.UpdateCart(input.CartID, input.Quantity)
	if !status {
		return c.JSON(400, helper.ResponseFormat("Failed", 400, "Not enough stock!", nil, nil))
	}
	if err != nil {
		return c.JSON(500, helper.ResponseFormat("Failed", 500, "Server Error!", nil, nil))
	}

	return c.JSON(200, helper.ResponseFormat("Succeed", 200, "Shopping Cart item quantity successfully updated!", nil, nil))
}

func (dth *DetailTransactionHandler) DeleteCart(c echo.Context) error {
	// Route: DELETE /cart?cart_id=1
	cartID, _ := strconv.Atoi(c.QueryParam("cart_id"))

	err := dth.srv.DeleteCart(uint(cartID))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat("Failed", 500, "Server Error!", nil, nil))
	}

	return c.JSON(200, helper.ResponseFormat("Succeed", 201, "Shopping Cart item successfully deleted!", nil, nil))
}
