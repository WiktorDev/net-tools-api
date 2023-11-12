package routes

import (
	"github.com/labstack/echo/v4"
	"net-tools-api/routes/dns"
)

func BootstrapRouting(echo *echo.Echo) {
	echo.GET("/dns/records/:domain", dns.FetchTxTRecords)
}
