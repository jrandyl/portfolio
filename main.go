package main

import (
	"database/sql"
	"log"

	"github.com/jrandyl/portfolio/database/sqlc"
	"github.com/jrandyl/portfolio/server"
	"github.com/jrandyl/portfolio/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configurations: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database:", err)
	}

	transaction := sqlc.NewTransaction(conn)
	server, err := server.NewServer(config, transaction)
	if err != nil {
		log.Fatal("cannot create the server", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start the server", err)
	}
}
