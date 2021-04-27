package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/juliandev/bookstore_users-api/domain/users"
	"fmt"
	"github.com/juliandev/bookstore_users-api/services"
	"github.com/juliandev/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implemented me!")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implemented me!")
}
