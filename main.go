package main

import (
	"github.com/mar-tina/smailtrail/smailclient"
)

var MySmailClient smailclient.ISmailClient

func main() {

	initSmailClient()

	// e := echo.New()

	// // Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// // Route => handler
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!\n")
	// })

	// e.Static("/static", "assets")

	// // Start server
	// e.Logger.Fatal(e.Start(":1323"))

}

func initSmailClient() {
	MySmailClient = &smailclient.SmailClient{}
	MySmailClient.NewSmailClient("credentials.json")
	MySmailClient.ListLabels()
	MySmailClient.ListMessages()
}
