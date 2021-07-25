package services

import (
	"context"

	"github.com/leonardonatali/full-cycle/grpc/pkg/entities"
	"github.com/leonardonatali/full-cycle/grpc/pkg/protobuf/users"
)

type UserService struct {
	users.UnimplementedUserServiceServer
	usersRepo entities.UsersRepository
}

func NewUsersService(usersRepo entities.UsersRepository) *UserService {
	return &UserService{
		usersRepo: usersRepo,
	}
}

func (s UserService) Add(ctx context.Context, req *users.User) (user *users.User, err error) {
	res, err := s.usersRepo.Add(&entities.User{
		Name:  req.Name,
		Email: req.Email,
	})

	if err != nil {
		return
	}

	return &users.User{
		Id:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}, nil
}
