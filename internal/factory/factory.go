package factory

import (
	"TokoGadget/configs"
	"TokoGadget/internal/features/users/handler"
	"TokoGadget/internal/features/users/repository"
	"TokoGadget/internal/features/users/services"
	"TokoGadget/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	cfg := configs.ImportSetting()
	db, _ := configs.ConnectDB(cfg)
	db.AutoMigrate(&repository.User{})
	um := repository.NewUserModel(db)
	us := services.NewUserService(um)
	uc := handler.NewUserController(us)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	// Register

	// t.GET("", tc.ShowMyTodo())
	// t.POST("", tc.CreateTodo())

	routes.InitRoute(e, uc)
}