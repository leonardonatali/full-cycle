package main

import (
	"github.com/leonardonatali/full-cycle/grpc/config"
	"github.com/leonardonatali/full-cycle/grpc/pkg/server"
)

func main() {
	cfg := &config.Config{
		GRPCPort: 50050,
	}

	srv := server.New(cfg)
	srv.Run()
}
