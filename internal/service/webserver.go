package service

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mar-tina/smailtrail/internal/smailclient"
)

var MySmailClient smailclient.ISmailClient

func StartWebServer(port string) {
	e := echo.New()

	// r := mux.NewRouter()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.File("/", "public/index.html")

	e.GET("/labels", ListLabelsHandler)
	e.GET("/initialauth", InitialAuth)
	e.Static("/*", "public/assets")

	e.Logger.Fatal(e.Start(":" + port))
}
