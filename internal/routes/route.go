package routes

import (
	"TokoGadget/internal/features/payment"
	"TokoGadget/internal/features/users"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)


func InitRoute(e *echo.Echo, uc users.UserHandler, pc payment.PaymentHandler) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
	e.PUT("/users", uc.Update, echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		}))
	e.GET("/users", uc.GetProfile, echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		}))
	e.DELETE("/users", uc.Delete, echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		}))
	// Rute untuk pembayaran dengan Midtrans Snap API
	e.POST("https://app.sandbox.midtrans.com/snap/v1/transactions", pc.CreateSnapTransaction, echojwt.WithConfig(
		echojwt.Config{
			SigningKey:    []byte("passkeyJWT"),
			SigningMethod: jwt.SigningMethodHS256.Name,
		}))
	
}
// func setRoute(e *echo.Echo) {
// 	t := e.Group("/users")
// 	t.Use(echojwt.WithConfig(
// 		echojwt.Config{
// 			SigningKey:    []byte("passkeyJWT"),
// 			SigningMethod: jwt.SigningMethodHS256.Name,
// 		},
// 	))
// }

