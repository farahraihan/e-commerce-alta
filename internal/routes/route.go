package routes

import (
	"TokoGadget/configs"
	dt_hnd "TokoGadget/internal/features/detail_transactions"
	p_hnd "TokoGadget/internal/features/products"
	s_hnd "TokoGadget/internal/features/sales"
	t_hnd "TokoGadget/internal/features/transactions"
	u_hnd "TokoGadget/internal/features/users"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uc u_hnd.Handler, th t_hnd.THandler, dth dt_hnd.DTHandler, ph p_hnd.PHandler, sh s_hnd.SHandler) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())

	UsersRoute(e, uc)
	TransactionsRoute(e, th, dth)
	ProductsRoute(e, ph)
	SalesRoute(e, sh)
}

func UsersRoute(e *echo.Echo, uc u_hnd.Handler) {
	u := e.Group("/users")
	u.Use(JWTConfig())
	u.PUT("", uc.Update)
	u.GET("", uc.GetProfile)
	u.DELETE("", uc.Delete)
}

func TransactionsRoute(e *echo.Echo, th t_hnd.THandler, dth dt_hnd.DTHandler) {
	c := e.Group("/cart")
	c.Use(JWTConfig())
	c.POST("", dth.AddToCart)
	c.GET("", dth.GetAllCart)
	c.PUT("", dth.UpdateCart)
	c.DELETE("", dth.DeleteCart)

	t := e.Group("/transaction")
	t.Use(JWTConfig())
	t.GET("", th.GetAllTransactions)
	t.PUT("/:transaction_id", th.Checkout)
	t.GET("/:transaction_id", th.GetTransaction)
	t.DELETE("/:transaction_id", th.DeleteTransaction)
	e.POST("/midtrans_update", th.CheckStatusPayment) //Midtrans Callback
}

func ProductsRoute(e *echo.Echo, ph p_hnd.PHandler) {
	p := e.Group("/products")
	p.GET("", ph.GetAllProducts())
	p.GET("/:product_id", ph.GetProductByID())
	p.POST("", ph.AddProduct(), JWTConfig())
	p.PUT("/:product_id", ph.UpdateProductByID(), JWTConfig())
	p.DELETE("/:product_id", ph.DeleteProduct(), JWTConfig())
}

func SalesRoute(e *echo.Echo, sh s_hnd.SHandler) {
	s := e.Group("/sales")
	s.GET("/user/:user_id", sh.GetSalesByUserID())
	s.GET("/:sales_id", sh.GetSalesByTransactionID(), JWTConfig())
}

func JWTConfig() echo.MiddlewareFunc {
	return echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte(configs.ImportPasskey()),
			SigningMethod: jwt.SigningMethodHS256.Name,
		},
	)
}
