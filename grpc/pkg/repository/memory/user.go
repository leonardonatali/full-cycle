package memory

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/leonardonatali/full-cycle/grpc/pkg/entities"
)

type UsersMemoryDB struct {
	data []*entities.User
}

func (db *UsersMemoryDB) Add(user *entities.User) (*entities.User, error) {
	user.ID = uuid.NewString()
	db.data = append(db.data, user)

	return user, nil
}

func (db *UsersMemoryDB) GetByID(id string) (*entities.User, error) {
	for _, user := range db.data {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user with id: %s not found", id)
}
