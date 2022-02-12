package services

import (
	"github.com/tfregonese/bookstore_users-api/domain/users"
	"github.com/tfregonese/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(user_id int64) (*users.User, *errors.RestErr) {
	user := users.User{
		Id: user_id,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}
