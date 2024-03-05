package main

import (
	"log"

	"github.com/jrandyl/portfolio/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatal("cannot create the server", err)
	}

	err = server.Start(":11000")

	if err != nil {
		log.Fatal("cannot start the server", err)
	}
}
