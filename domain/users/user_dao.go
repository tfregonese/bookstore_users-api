package users

import (
	"fmt"
	"github.com/tfregonese/bookstore_users-api/datasources/mysql/users_db"
	"github.com/tfregonese/bookstore_users-api/logger"
	"github.com/tfregonese/bookstore_users-api/utils/error_utils"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status, password FROM users WHERE status=?;"
)

var (
	usersDB = users_db.Client
)

func (user *User) Get() *error_utils.RestErr {
	stmt, err := usersDB.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return error_utils.NewInternalServerError("database_error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to execute get user statement", err)
		return error_utils.NewInternalServerError("database_error")
	}

	return nil
}

func (user *User) Save() *error_utils.RestErr {

	// -- you can replace this query using insertResult, err := stmt.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	stmt, err := usersDB.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return error_utils.NewInternalServerError("database_error")
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error("error when trying to execute save user statement", err)
		return error_utils.NewInternalServerError("database_error")
	}
	// --

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to execute save user statement", err)
		return error_utils.NewInternalServerError("database_error")
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *error_utils.RestErr {
	stmt, err := usersDB.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return error_utils.NewInternalServerError("database_error")
	}
	defer stmt.Close()

	if restErr := user.Validate(); restErr != nil {
		return restErr
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id, user.Status, user.Password)
	if err != nil {
		logger.Error("error when trying to execute update user statement", err)
		return error_utils.NewInternalServerError("database_error")
	}

	return nil
}

func (user *User) Delete() *error_utils.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return error_utils.NewInternalServerError("database_error")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to execute delete user statement", err)
		return error_utils.NewInternalServerError("database_error")
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *error_utils.RestErr) {

	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, error_utils.NewInternalServerError("database_error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to execute find user by status statement", err)
		return nil, error_utils.NewInternalServerError("database_error")
	}
	defer rows.Close()

	results := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
			logger.Error("error when trying to execute find user by status statement", err)
			return nil, error_utils.NewInternalServerError("database_error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, error_utils.NewNotFoundError(fmt.Sprintf("no users with the status %s.", status))
	}

	return results, nil
}
