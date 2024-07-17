package routes

import (
	"TokoGadget/configs"
	dt_hnd "TokoGadget/internal/features/detail_transactions"
	t_hnd "TokoGadget/internal/features/transactions"
	u_hnd "TokoGadget/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uc u_hnd.Handler, th t_hnd.THandler, dth dt_hnd.DTHandler) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())

	TransactionsRoute(e, th, dth)
}

func TransactionsRoute(e *echo.Echo, th t_hnd.THandler, dth dt_hnd.DTHandler) {
	c := e.Group("/cart")
	c.Use(JWTConfig())
	c.POST("", dth.AddToCart)
	c.GET("", dth.GetAllCart)
	c.PUT("", dth.UpdateCart)
	c.DELETE("/:cart_id", dth.DeleteCart)

	t := e.Group("/transaction")
	t.Use(JWTConfig())
	t.GET("", th.GetAllTransactions)
	t.PUT("/:transaction_id", th.Checkout)
	t.GET("/:transaction_id", th.GetTransaction)
	t.DELETE("/:transaction_id", th.DeleteTransaction)
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(configs.ImportPasskey()),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
