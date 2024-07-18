package payment

import (
	"github.com/labstack/echo/v4"
)


type PaymentHandler interface {
	CreateSnapTransaction(c echo.Context) error
}