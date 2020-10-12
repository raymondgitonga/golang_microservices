package dormain

import (
	"fmt"
	"github.com/raymondgitonga/golang_microservices/mvc/utils"
	"log"
	"net/http"
)

type userDao struct{}

type userDaoInterface interface {
	GetUser(userId int64) (*User, *utils.AppError)
}

var (
	users = map[int64]*User{
		123: &User{123, "Raymond", "Gitonga", "raytosh95@gmail.com"},
	}

	UserDaoInterface userDaoInterface
)

func init() {
	UserDaoInterface = &userDao{}
}

func (u *userDao) GetUser(userId int64) (*User, *utils.AppError) {
	log.Println("We are accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.AppError{
		Message:    fmt.Sprintf("User %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       utils.UserNotFound,
	}
}
