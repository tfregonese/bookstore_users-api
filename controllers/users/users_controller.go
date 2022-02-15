package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tfregonese/bookstore_users-api/domain/users"
	"github.com/tfregonese/bookstore_users-api/services"
	"github.com/tfregonese/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id.")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		//Handle error
		restErr := errors.NewBadRequestError("Invalid Json")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//Handle saveErr
		c.JSON(http.StatusBadRequest, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user id.")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json.")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch
	user.Id = userId
	result, saveErr := services.UpdateUser(isPartial, user)
	if saveErr != nil {
		//Handle saveErr
		c.JSON(http.StatusBadRequest, saveErr)
		return
	}

	c.JSON(http.StatusOK, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
