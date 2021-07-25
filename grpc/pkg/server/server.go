package server

import (
	"fmt"
	"log"
	"net"

	"github.com/leonardonatali/full-cycle/grpc/config"
	"github.com/leonardonatali/full-cycle/grpc/pkg/protobuf/users"
	"github.com/leonardonatali/full-cycle/grpc/pkg/repository/memory"
	"github.com/leonardonatali/full-cycle/grpc/pkg/services"
	"google.golang.org/grpc"
)

type Server struct {
	cfg *config.Config
}

func New(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GRPCPort))
	if err != nil {
		log.Fatalf("could not start the listener: %s", err.Error())
	}

	srv := grpc.NewServer()

	repo := memory.UsersMemoryDB{}

	userService := services.NewUsersService(&repo)
	users.RegisterUserServiceServer(srv, userService)

	if err := srv.Serve(listener); err != nil {
		log.Fatalf("could not start gRPC server: %s", err.Error())
	}
}
