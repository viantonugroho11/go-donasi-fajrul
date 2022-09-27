package http

import (
"github.com/labstack/echo"
donasiSrv "donasi/service"
)

func NewDonasiHandler(app *echo.Group, donasiSrv *donasiSrv.DonasiService) {
	app.GET("/", GetAllDonasi(*donasiSrv))
	app.GET("/:id", GetDetailDonasi(*donasiSrv))
}

func GetAllDonasi(donasiSrv donasiSrv.DonasiService) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "OK")
	}
}

func GetDetailDonasi(donasiSrv donasiSrv.DonasiService) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "OK")
	}
}