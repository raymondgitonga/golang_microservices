package service

import (
	"github.com/raymondgitonga/golang_microservices/mvc/dormain"
	"github.com/raymondgitonga/golang_microservices/mvc/utils"
)

type userService struct{}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*dormain.User, *utils.AppError) {
	user, err := dormain.UserDao.GetUser(userId)

	if err != nil {
		return nil, err
	}
	return user, nil
}
