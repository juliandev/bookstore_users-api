package app

import (
	"github.com/juliandev/bookstore_users-api/controllers/ping"
	"github.com/juliandev/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/search_id", users.SearchUser)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
}
