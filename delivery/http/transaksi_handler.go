package http

import (
	"donasi/common"
	"donasi/model"
	transSrv "donasi/service"

	"github.com/labstack/echo"
)

func NewTransaksiHandler(app *echo.Group, transSrv *transSrv.TransaksiService) {
	app.POST("/", PostTranscationDonasi(*transSrv))
	// app.GET("/notification-handler", GetNotificationHandler(*transSrv))

}

func PostTranscationDonasi(transSrv transSrv.TransaksiService) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqPayload := &model.PayloadTransaksiRequest{}
		if err := c.Bind(reqPayload); err != nil {
			return c.JSON(400, err.Error())
		}
		if err := common.ValidateStruct(reqPayload); err != nil {
			return c.JSON(400, err)
		}

		result ,err := transSrv.PostTranscationDonasi(c.Request().Context(), reqPayload)
		if err != nil {
			return c.JSON(400, err)
		}
		return c.JSON(200, result)
	}
}


func GetNotificationHandler(transSrv *transSrv.TransaksiService) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "OK")
	}
}
