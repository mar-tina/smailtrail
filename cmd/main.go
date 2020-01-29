package main

import (
	"github.com/mar-tina/smailtrail/internal/dbclient"
	"github.com/mar-tina/smailtrail/internal/service"
	"github.com/mar-tina/smailtrail/internal/smailclient"
)

func main() {

	initSmailClient()
	initDatabase()

	service.StartWebServer("8000")
}

func initSmailClient() {
	service.MySmailClient = &smailclient.SmailClient{}
}

func initDatabase() {
	smailclient.DBClient = &dbclient.BadgerClient{}
	smailclient.DBClient.OpenBadgerDB("./storage")
}
