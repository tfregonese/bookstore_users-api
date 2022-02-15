package errors

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func HandleError(err error) *RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		return NewInternalServerError(err.Error())
	}

	switch sqlErr.Number {
	case 1062:
		return NewBadRequestError(sqlErr.Message)
	case 1054:
		return NewInternalServerError(sqlErr.Message)
	default:
		fmt.Println(sqlErr.Number)
		fmt.Println(sqlErr.Message)
		return NewInternalServerError(err.Error())
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
		Message: message,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusNotFound,
		Error:   "not_found",
		Message: message,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
		Message: message,
	}
}
