package main

import (
	"database/sql"
	"github.com/golang/simplebank/api"
	db "github.com/golang/simplebank/db/sqlc"
	"github.com/golang/simplebank/gapi"
	"github.com/golang/simplebank/pb"
	"github.com/golang/simplebank/util"
	_ "github.com/lib/pq"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	//server, err := api.NewServer(config, store)
	//if err != nil {
	//	log.Fatal("cannot create server:", err)
	//}
	//
	//err = server.Start(config.ServerAddress)
	//if err != nil {
	//	log.Fatal("Cannot connect to db:", err)
	//}
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	grpcServer := grpc2.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
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

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
}
