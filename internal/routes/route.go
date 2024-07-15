package routes

import (
	"TokoGadget/internal/features/users"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo, uc users.Handler) {
	e.POST("/register", uc.Register())
	e.POST("/login", uc.Login())
}

