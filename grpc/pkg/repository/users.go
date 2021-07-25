package repository

import "github.com/leonardonatali/full-cycle/grpc/pkg/entities"

type UsersRepository interface {
	GetByID(id string) (*entities.User, error)
	Add(user *entities.User) (*entities.User, error)
}
