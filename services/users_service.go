package services

import (
	"github.com/tfregonese/bookstore_users-api/domain/users"
	"github.com/tfregonese/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

func GetUser(user_id int64) (*users.User, *errors.RestErr) {
	return &users.User{
		Id:        123,
		FirstName: "Pepe",
		LastName:  "Pepitos",
	}, nil
}
