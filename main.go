package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sheyi103/agtMiddleware/api"
	db "github.com/sheyi103/agtMiddleware/db/sqlc"
)

const (
	dbDriver = "mysql"
	dbSource = "agt:Password123@tcp(localhost:3306)/agt_middleware_db?parseTime=true"

	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
