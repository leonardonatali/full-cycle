package entities

type UsersRepository interface {
	GetByID(id string) (*User, error)
	Add(user *User) (*User, error)
}

type User struct {
	ID    string
	Name  string
	Email string
}
