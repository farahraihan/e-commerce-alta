package routes

import (
	"TokoGadget/internal/features/products"
	"TokoGadget/internal/features/users"

	"github.com/golang-jwt/jwt"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uc users.Handler, pc products.PHandler) {
	// Routes for users
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

	// Routes for products
	productGroup := e.Group("/products")
	productGroup.Use(echojwt.JWT([]byte("passkeyJWT"))) // Middleware JWT untuk produk

	productGroup.POST("", pc.AddProduct())
	productGroup.GET("", pc.GetAllProducts())
	productGroup.GET("/:product_id", pc.GetProductByID())
	productGroup.PUT("/:product_id", pc.UpdateProductByID())
	productGroup.DELETE("/:product_id", pc.DeleteProduct())
}
