package user

import (
	"github.com/gin-gonic/gin"
	"github.com/nursace/bookstore_users-api/domain/users"
	"github.com/nursace/bookstore_users-api/services"
	"github.com/nursace/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewRequestBodyError("invalid json body")
		c.JSON(restErr.Status, restErr)
		// TODO: Return BadRequest to a caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TODO: handle user creation error

		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
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
