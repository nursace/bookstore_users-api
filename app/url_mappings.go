package app

import (
	"github.com/nursace/bookstore_users-api/controller/ping"
	"github.com/nursace/bookstore_users-api/controller/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", user.GetUser)
	router.POST("/users", user.CreateUser)
}
