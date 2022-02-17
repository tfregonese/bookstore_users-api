package users

import (
	"fmt"
	"github.com/tfregonese/bookstore_users-api/datasources/mysql/users_db"
	"github.com/tfregonese/bookstore_users-api/utils/date_utils"
	"github.com/tfregonese/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

var (
	usersDB = users_db.Client
)

func (user *User) Get() *errors.RestErr {
	stmt, err := usersDB.Prepare(queryGetUser)
	if err != nil {
		return errors.HandleError(err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)
		return errors.HandleError(err)
	}

	return nil
}

func (user *User) Save() *errors.RestErr {

	// -- you can replace this query using insertResult, err := stmt.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	stmt, err := usersDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.HandleError(err)
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.HandleError(err)
	}
	// --

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.HandleError(err)
	}
	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := usersDB.Prepare(queryUpdateUser)
	if err != nil {
		return errors.HandleError(err)
	}
	defer stmt.Close()

	if restErr := user.Validate(); restErr != nil {
		return restErr
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.HandleError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.HandleError(err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return errors.HandleError(err)
	}

	return nil
}
