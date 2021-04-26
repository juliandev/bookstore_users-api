package services

import (
	"github.com/juliandev/bookstore_users-api/domain/users"
)

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
