package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/internal/api"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret123@localhost:5432/eloc_auth?sslmode=disable"
	serverAddress = "0.0.0.0:8081"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	store := db.NewStore(conn)
	runGinServer(store)
}

func runGinServer (store db.Store) {
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create sever:", err)
	}
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
