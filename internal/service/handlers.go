package service

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/labstack/echo"
	"github.com/mar-tina/smailtrail/internal/auth"
	"github.com/mar-tina/smailtrail/internal/smailclient"
)

var MySmailClient smailclient.ISmailClient

type AuthCode struct {
	code string
}

func ListLabelsHandler(c echo.Context) error {
	log.Println("Im i even here")
	res, err := MySmailClient.ListLabels()
	if err != nil {
		log.Println("Failed to fetch labels", err.Error())
	}
	log.Println("The res", res)
	jsonBytes, _ := json.Marshal(res)
	return c.String(200, string(jsonBytes))
}

func ListAllMessages(c echo.Context) error {

	tokparam := c.QueryParam("nextpagetoken")

	list, err := MySmailClient.ListMessages(tokparam)
	if err != nil {
		return c.String(500, err.Error())
	}

	var res map[string]interface{}
	res = make(map[string]interface{}, 2)

	res["list"] = list
	res["msgs"] = msgs
	return c.JSON(200, res)
}

func ListAllSubscriptions(c echo.Context) error {
	take, err := strconv.Atoi(c.QueryParam("take"))
	skip, err := strconv.Atoi(c.QueryParam("skip"))
	if err != nil {
		c.JSON(500, "Something went wrong")
	}

	subs, err := smailclient.StormDBClient.FetchSubscriptions(take, skip)
	if err != nil {
		c.JSON(500, "Something went wrong")
	}

	return c.JSON(200, subs)
}

func InitialAuth(c echo.Context) error {
	authURL, err := auth.GetAuthURL("credentials.json")
	message := struct {
		Message string
		err     error
	}{
		"Something went wrong. Make sure you have the credentials.json file",
		err,
	}
	if err != nil {
		return c.JSON(500, message)
	}

	jsonBytes, _ := json.Marshal(authURL)
	return c.String(200, string(jsonBytes))
}

func ProcessToken(c echo.Context) error {
	m := echo.Map{}

	err := c.Bind(&m)

	if err != nil {
		return c.String(500, err.Error())
	}

	err = MySmailClient.InitSmailClient("credentials.json", m["code"].(string))
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.JSON(200, "You have been successfully authenticated")
}
