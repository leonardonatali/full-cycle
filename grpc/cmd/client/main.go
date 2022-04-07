package main

import (
	"context"
	"fmt"
	"io"

	"github.com/leonardonatali/full-cycle/grpc/config"
	"github.com/leonardonatali/full-cycle/grpc/pkg/protobuf/users"
	"google.golang.org/grpc"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(fmt.Sprintf("0.0.0.0:%d", cfg.GRPCPort), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := users.NewUserServiceClient(conn)

	res, err := client.AddVerbose(context.Background(), &users.User{
		Name:  "Leonardo",
		Email: "test@mail.com",
	})

	if err != nil {
		panic(err)
	}

	for {
		stream, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		fmt.Printf("status: %s\nuser: %#v\n\n", stream.Status, stream.GetUser())
	}
}
