package services

import (
	"github.com/tfregonese/bookstore_users-api/domain/users"
	"github.com/tfregonese/bookstore_users-api/utils/crypto_utils"
	"github.com/tfregonese/bookstore_users-api/utils/date_utils"
	"github.com/tfregonese/bookstore_users-api/utils/error_utils"
)

var (
	UserService usersServiceInterface = &userService{}
)

type usersServiceInterface interface {
	Create(users.User) (*users.User, *error_utils.RestErr)
	Get(int64) (*users.User, *error_utils.RestErr)
	Update(bool, users.User) (*users.User, *error_utils.RestErr)
	Delete(int64) *error_utils.RestErr
	Search(string) (users.Users, *error_utils.RestErr)
	LogInUser(users.LoginRequest) (*users.User, *error_utils.RestErr)
}

type userService struct{}

func (s *userService) Create(user users.User) (*users.User, *error_utils.RestErr) {

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

func (s *userService) Get(userId int64) (*users.User, *error_utils.RestErr) {

	user := users.User{
		Id: userId,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userService) Update(isPartial bool, user users.User) (*users.User, *error_utils.RestErr) {

	current, err := s.Get(user.Id)
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

func (s *userService) Delete(userId int64) *error_utils.RestErr {

	user := &users.User{Id: userId}

	return user.Delete()
}

func (s *userService) Search(userStatus string) (users.Users, *error_utils.RestErr) {

	user := &users.User{}

	return user.FindByStatus(userStatus)
}

func (s *userService) LogInUser(request users.LoginRequest) (*users.User, *error_utils.RestErr) {

	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}

	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return dao, nil
}
