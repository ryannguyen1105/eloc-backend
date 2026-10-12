package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	pb "github.com/ryannguyen1105/eloc-backend/api/pb/eloc_auth"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/config"
	db "github.com/ryannguyen1105/eloc-backend/services/eloc_auth/db/sqlc"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/gapi"
	"github.com/ryannguyen1105/eloc-backend/services/eloc_auth/internal/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	// runGrpcServer(config, store)
	runGinServer(config, store)
}

func runGrpcServer(config config.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterEloc_AuthServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func runGinServer(config config.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create sever:", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
