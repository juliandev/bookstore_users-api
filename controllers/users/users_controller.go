package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/juliandev/bookstore_users-api/domain/users"
	"fmt"
	"github.com/juliandev/bookstore_users-api/services"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Handle json error
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle error
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
