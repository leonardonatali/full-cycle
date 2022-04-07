package services

import (
	"context"
	"io"
	"log"
	"time"

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
		Name:  req.GetName(),
		Email: req.GetEmail(),
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

func (s UserService) AddVerbose(req *users.User, stream users.UserService_AddVerboseServer) error {
	var err error
	err = stream.Send(&users.UserResultStream{
		Status: "Initializing",
		User:   &users.User{},
	})

	if err != nil {
		return err
	}

	time.Sleep(3 * time.Second)

	res, err := s.usersRepo.Add(&entities.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	})

	if err != nil {
		return err
	}

	user := &users.User{
		Id:    res.ID,
		Name:  res.Name,
		Email: res.Email,
	}

	err = stream.Send(&users.UserResultStream{
		Status: "Created",
		User:   user,
	})

	if err != nil {
		return err
	}

	time.Sleep(3 * time.Second)

	err = stream.Send(&users.UserResultStream{
		Status: "Finished",
		User:   user,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s UserService) AddUsers(stream users.UserService_AddUsersServer) error {
	results := &users.Users{}

	for {
		received, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(results)
		}
		if err != nil {
			return err
		}

		log.Printf("New user received: %s", received.GetName())

		new, err := s.usersRepo.Add(&entities.User{
			Name: received.GetName(),
			ID:   received.GetEmail(),
		})

		if err != nil {
			return err
		}

		log.Printf("New user Name: %s\n", new.Name)
		log.Printf("New user ID: %s\n", new.ID)

		results.Users = append(results.Users, &users.User{
			Id:    new.ID,
			Name:  new.Name,
			Email: new.Email,
		})
	}
}

func (s UserService) AddBidirectionalUsers(stream users.UserService_AddBidirectionalUsersServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		if err := stream.Send(&users.UserResultStream{Status: "Initializing", User: &users.User{}}); err != nil {
			return err
		}

		if err := stream.Send(&users.UserResultStream{Status: "Creating", User: &users.User{}}); err != nil {
			return err
		}

		new, err := s.usersRepo.Add(&entities.User{
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})

		if err != nil {
			return err
		}

		user := &users.User{
			Id:    new.ID,
			Name:  new.Name,
			Email: new.Email,
		}

		if err := stream.Send(&users.UserResultStream{Status: "Created", User: user}); err != nil {
			return err
		}
	}
}
