package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/config"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/internal/api"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	runGinServer(config, store)
}

func runGinServer(config config.Config, store db.Store) {
	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create sever:", err)
	}
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
