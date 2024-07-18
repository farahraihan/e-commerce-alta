package factory

import (
	"TokoGadget/configs"

	// users
	"TokoGadget/internal/features/users/handler"
	"TokoGadget/internal/features/users/repository"
	"TokoGadget/internal/features/users/services"

	//Product
	// "TokoGadget/internal/features/products/handler"
	product "TokoGadget/internal/features/products/repository"
	// "TokoGadget/internal/features/products/services"

	// transaction
	transaction "TokoGadget/internal/features/transactions/repository"

	// detail transaction
	detailTransaction "TokoGadget/internal/features/detail_transactions/repository"

	"TokoGadget/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	cfg := configs.ImportSetting()
	db, _ := configs.ConnectDB(cfg)
	db.AutoMigrate(&repository.User{}, &product.Product{}, &transaction.Transaction{}, &detailTransaction.DetailTransaction{})
	um := repository.NewUserModel(db)
	us := services.NewUserService(um)
	uc := handler.NewUserController(us)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	// e.AutoTLSManager.Cache = autocert.DirCache("/var/www/.cache")
	// e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	// e.GET("/", func(c echo.Context) error {
	// 	return c.HTML(http.StatusOK, `
	// 		<h1>Welcome to Echo!</h1>
	// 		<h3>TLS certificates automatically installed from Let's Encrypt :)</h3>
	// 	`)
	// })

	routes.InitRoute(e, uc)
}
