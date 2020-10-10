package app

import (
	"github.com/raymondgitonga/golang_microservices/mvc/controller"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controller.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
