package main

import (
	conf "donasi/config"
httpHandler "donasi/delivery/http"
"donasi/repository"
"donasi/service"
	"github.com/labstack/echo"
)

func main() {
	// Write your code here
	config := conf.New()

	confDb := conf.InitDB(config)
	confMidtrans := conf.InitMidtrans(config)

	dbRepo := repository.NewDatabaseRepository(confDb,confMidtrans)

	dnsService := service.NewDonasiService(dbRepo)
	e := echo.New()
	api := e.Group("/api")
	e.GET("/health",func(c echo.Context) error {
		return c.JSON(200, "OK")
	})

	httpHandler.NewDonasiHandler(api.Group("/v1/donasi"), &dnsService)

	e.Logger.Fatal(e.Start(":1324"))
}