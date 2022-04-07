package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	addVerbose(client)
	addUsers(client)
}

func addVerbose(client users.UserServiceClient) {
	res, err := client.AddVerbose(context.Background(), &users.User{
		Name:  "Leonardo",
		Email: "test@mail.com",
	})

	if err != nil {
		log.Fatalf("could not add user: %s", err.Error())
	}

	for {
		stream, err := res.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("could receive add user stream: %s", err.Error())
		}

		log.Printf("status: %s\nuser: %#v\n\n", stream.Status, stream.GetUser())
	}
}

func addUsers(client users.UserServiceClient) {
	items := []*users.User{
		{
			Name:  "Leonardo",
			Email: "test@mail.com",
		},
		{
			Name:  "João",
			Email: "test2@mail.com",
		},
	}

	c, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("could not get add users client: %s", err.Error())
	}

	for _, item := range items {
		log.Printf("sending user: %s", item.Name)

		if err := c.Send(item); err != nil {
			log.Fatalf("could not send user: %s", err.Error())
		}

		time.Sleep(time.Second)
	}

	res, err := c.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not close and receive add users stream: %s", err.Error())
	}

	for _, user := range res.GetUsers() {
		log.Printf("-----\ninserted user ID: %s\ninserted user name: %s", user.GetId(), user.GetName())
	}
}
