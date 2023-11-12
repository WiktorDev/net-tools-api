package dns

import (
	"github.com/labstack/echo/v4"
	"net"
)

func FetchTxTRecords(c echo.Context) error {
	records, err := net.LookupTXT(c.Param("domain"))
	if err != nil {
		return c.JSON(400, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"records": records,
	})
}
