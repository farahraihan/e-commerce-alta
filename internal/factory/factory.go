package factory

import (
	"TokoGadget/configs"

	// users
	"TokoGadget/internal/features/users/handler"
	"TokoGadget/internal/features/users/repository"
	"TokoGadget/internal/features/users/services"

	//Product
	productHandler "TokoGadget/internal/features/products/handler"
	productRepository "TokoGadget/internal/features/products/repository"
	productServices "TokoGadget/internal/features/products/services"

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
	db.AutoMigrate(&repository.User{}, &productRepository.Product{}, &transaction.Transaction{}, &detailTransaction.DetailTransaction{})
	um := repository.NewUserModel(db)
	us := services.NewUserService(um)
	uc := handler.NewUserController(us)

	pm := productRepository.NewProductModel(db)
	ps := productServices.NewProductService(pm)
	pc := productHandler.NewProductController(ps)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	// Register

	// t.GET("", tc.ShowMyTodo())
	// t.POST("", tc.CreateTodo())

	routes.InitRoute(e, uc, pc)
}
