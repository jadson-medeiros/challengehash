package main

import (
	"github.com/jadson-medeiros/challengehash/server"
)

func main() {
	server := server.NewServer()

	server.Run()
}
