package main

import (
	"github.com/jadson-medeiros/challengehash/database"
	"github.com/jadson-medeiros/challengehash/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()

	server.Run()
}
