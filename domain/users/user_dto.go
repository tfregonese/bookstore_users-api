package users

import (
	"strings"

	"github.com/tfregonese/bookstore_users-api/utils/error_utils"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

func (user User) Validate() *error_utils.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return error_utils.NewNotFoundError("Invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return error_utils.NewNotFoundError("Invalid password")
	}
	return nil
}
