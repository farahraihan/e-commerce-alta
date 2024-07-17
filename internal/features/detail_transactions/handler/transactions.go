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

func (dth *DetailTransactionHandler) AddToCart(c echo.Context) error {
	// Route: POST /cart
	userID := dth.tu.DecodeToken(c.Get("user").(*jwt.Token))

	var productID uint
	err := c.Bind(&productID)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "Input Error!", nil))
	}

	err = dth.srv.AddToCart(productID, userID)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(201, helper.ResponseFormat(201, "Shopping Cart datas successfully retreived!", nil))
}

func (dth *DetailTransactionHandler) GetAllCart(c echo.Context) error {
	// Route: GET /cart
	userID := dth.tu.DecodeToken(c.Get("user").(*jwt.Token))

	result, err := dth.srv.GetAllCart(uint(userID))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helper.ResponseFormat(200, "Shopping Cart datas successfully retreived!", result))
}

func (dth *DetailTransactionHandler) UpdateCart(c echo.Context) error {
	// Route: PUT /cart
	var input RequestCart
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "Input Error!", nil))
	}

	err = dth.srv.UpdateCart(input.CartID, input.Quantity)
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helper.ResponseFormat(201, "Shopping Cart item quantity successfully updated!", nil))
}

func (dth *DetailTransactionHandler) DeleteCart(c echo.Context) error {
	// Route: PUT /cart
	cartID, _ := strconv.Atoi(c.Param("cart_id"))

	err := dth.srv.DeleteCart(uint(cartID))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "Server Error!", nil))
	}

	return c.JSON(200, helper.ResponseFormat(201, "Shopping Cart item successfully deleted!", nil))
}
