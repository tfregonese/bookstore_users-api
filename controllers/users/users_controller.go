package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tfregonese/bookstore_users-api/domain/users"
	"github.com/tfregonese/bookstore_users-api/services"
	"github.com/tfregonese/bookstore_users-api/utils/errors"
)

func GetUser(c *gin.Context) {
	user, err := services.GetUser(1)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User
	/*
			bytes, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				// Handle the Error
				c.String(http.StatusBadRequest, err.Error())
				return
			}
			if err := json.Unmarshal(bytes, &user); err != nil {
				// If there is no error in the body
				c.String(http.StatusBadRequest, err.Error())
				return
			}
		This can be replaced for:
	*/
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

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
