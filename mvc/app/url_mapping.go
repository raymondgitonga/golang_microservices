package app

import (
	"github.com/raymondgitonga/golang_microservices/mvc/controller"
)

func mapUrls() {
	router.GET("/users/:user_id", controller.GetUser)
}
