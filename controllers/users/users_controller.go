package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tfregonese/bookstore_users-api/domain/users"
	"github.com/tfregonese/bookstore_users-api/services"
	"github.com/tfregonese/bookstore_users-api/utils/error_utils"
)

func Get(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := error_utils.NewBadRequestError("Invalid user id.")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.UserService.Get(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//Handle error
		restErr := error_utils.NewBadRequestError("Invalid Json")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	result, saveErr := services.UserService.Create(user)
	if saveErr != nil {
		//Handle saveErr
		c.JSON(http.StatusBadRequest, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := error_utils.NewBadRequestError("Invalid user id.")
		c.JSON(err.Status, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := error_utils.NewBadRequestError("Invalid Json.")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	isPartial := c.Request.Method == http.MethodPatch
	user.Id = userId
	result, saveErr := services.UserService.Update(isPartial, user)
	if saveErr != nil {
		//Handle saveErr
		c.JSON(http.StatusBadRequest, saveErr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {

	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := error_utils.NewBadRequestError("Invalid user id.")
		c.JSON(err.Status, err)
		return
	}

	deleteErr := services.UserService.Delete(userId)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {

	userStatus := c.Query("user_status")
	if len(userStatus) == 0 {
		err := error_utils.NewBadRequestError("Invalid parameter.")
		c.JSON(err.Status, err)
		return
	}

	users, searchErr := services.UserService.Search(userStatus)
	if searchErr != nil {
		c.JSON(searchErr.Status, searchErr)
		return
	}

	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}
