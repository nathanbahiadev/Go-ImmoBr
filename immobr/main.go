package main

import (
	"github.com/nathanbahiadev/go-immobr/infrastructure/database"
	"github.com/nathanbahiadev/go-immobr/infrastructure/routes"
)

func main() {
	if _, err := database.ConnectToDatabase(); err != nil {
		panic(err.Error())
	}

	routes.StartServer()
}
