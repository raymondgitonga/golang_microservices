package service

import (
	"github.com/raymondgitonga/golang_microservices/mvc/dormain"
	"github.com/raymondgitonga/golang_microservices/mvc/utils"
)

func GetUser(userId int64) (*dormain.User, *utils.AppError) {
	return dormain.GetUser(userId)
}
