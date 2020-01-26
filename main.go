package main

import (
	"github.com/mar-tina/smailtrail/smailclient"
)

var MySmailClient smailclient.ISmailClient

func main() {
	initSmailClient()
}

func initSmailClient() {
	MySmailClient = &smailclient.SmailClient{}
	MySmailClient.NewSmailClient("credentials.json")
	MySmailClient.ListLabels()
}
