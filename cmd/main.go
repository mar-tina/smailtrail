package main

import (
	"github.com/mar-tina/smailtrail/internal/service"
	"github.com/mar-tina/smailtrail/internal/smailclient"
)

func main() {

	initSmailClient()

	service.StartWebServer("8000")
}

func initSmailClient() {
	service.MySmailClient = &smailclient.SmailClient{}
}
