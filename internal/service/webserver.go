package service

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartWebServer(port string) {
	e := echo.New()

	// Middleware

	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Logger.SetLevel(99)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.GET("/labels", ListLabelsHandler)
	e.GET("/initialauth", InitialAuth)
	e.GET("/allmessages", ListAllMessages)
	e.GET("/indietrail", GetIndividualTrail)
	e.POST("/completeauth", ProcessToken)

	e.Logger.Fatal(e.Start(":" + port))
}
