package main

import (
	"TokoGadget/internal/factory"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	factory.InitFactory(e)

	e.Logger.Error(e.StartAutoTLS(":8000"))
}
