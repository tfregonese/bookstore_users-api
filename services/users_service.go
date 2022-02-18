package services

import (
	"github.com/tfregonese/bookstore_users-api/domain/users"
	"github.com/tfregonese/bookstore_users-api/utils/crypto_utils"
	"github.com/tfregonese/bookstore_users-api/utils/date_utils"
	"github.com/tfregonese/bookstore_users-api/utils/error_utils"
)

func CreateUser(user users.User) (*users.User, *error_utils.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *error_utils.RestErr) {
	user := users.User{
		Id: userId,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *error_utils.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.FirstName = user.FirstName
		}
		if user.Email != "" {
			current.FirstName = user.FirstName
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(userId int64) *error_utils.RestErr {
	user := &users.User{Id: userId}

	return user.Delete()
}

func SearchUser(userStatus string) (users.Users, *error_utils.RestErr) {
	user := &users.User{}

	return user.FindByStatus(userStatus)
}
