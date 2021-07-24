package main

import (
	"log"

	"github.com/leonardonatali/full-cycle/grpc/config"
	"github.com/leonardonatali/full-cycle/grpc/pkg/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("cannot read configuration: %s", err.Error())
	}

	srv := server.New(cfg)
	srv.Run()
}
