package main

import (
	"github.com/labstack/echo/v4"
	"net-tools-api/routes"
)

func main() {
	e := echo.New()
	routes.BootstrapRouting(e)
	e.Logger.Fatal(e.Start(":1323"))
}
