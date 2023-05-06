package gapi

import (
	"fmt"
	db "github.com/golang/simplebank/db/sqlc"
	"github.com/golang/simplebank/pb"
	"github.com/golang/simplebank/token"
	"github.com/golang/simplebank/util"
)

// Server servers gRPC request for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer create a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
