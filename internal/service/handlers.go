package service

import (
	"encoding/json"

	"github.com/labstack/echo"
	"github.com/mar-tina/smailtrail/internal/auth"
)

func ListLabelsHandler(c echo.Context) error {
	res, _ := MySmailClient.ListLabels()
	jsonBytes, _ := json.Marshal(res)
	return c.String(200, string(jsonBytes))
}

func InitialAuth(c echo.Context) error {
	authURL := auth.GetAuthURL("credentials.json")
	jsonBytes, _ := json.Marshal(authURL)
	return c.String(200, string(jsonBytes))
}

func ProcessToken(c echo.Context) error {
	authcode := c.Param("code")
	srv := auth.GetSrvFromAuthCode(authcode)
	MySmailClient.InitSmailClient(srv)
	return c.String(201, "Auth was successful")
}
