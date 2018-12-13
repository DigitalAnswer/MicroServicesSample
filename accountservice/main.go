package main

import (
	"fmt"

	"github.com/DigitalAnswer/MicroServicesSample/accountservice/dbclient"

	"github.com/DigitalAnswer/MicroServicesSample/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	service.StartWebServer("8080")
}

func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDB()
	service.DBClient.Seed()
}
