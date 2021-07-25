package services

import (
	"context"

	"github.com/leonardonatali/full-cycle/grpc/pkg/protobuf/users"
)

type UserService struct {
	users.UnimplementedUserServiceServer
}

func (s UserService) Add(ctx context.Context, req *users.User) (user *users.User, err error) {
	return &users.User{}, nil
}
