package main

import (
	"github.com/mar-tina/smailtrail/internal/dbclient"
	"github.com/mar-tina/smailtrail/internal/service"
	"github.com/mar-tina/smailtrail/internal/smailclient"
)

func main() {

	initSmailClient()
	// initDatabase()
	initStormDB()

	service.StartWebServer("8000")
}

func initSmailClient() {
	service.MySmailClient = &smailclient.SmailClient{}
}

func initDatabase() {
	smailclient.DBClient = &dbclient.BadgerClient{}
	smailclient.DBClient.OpenBadgerDB("./storage")
}

func initStormDB() {
	smailclient.StormDBClient = &dbclient.StormClient{}
	smailclient.StormDBClient.OpenStormDB("smailtrail.db")
}
