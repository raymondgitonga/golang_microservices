package dormain

import (
	"fmt"
	"github.com/raymondgitonga/golang_microservices/mvc/utils"
	"log"
	"net/http"
)

var (
	users = map[int64]*User{
		123: &User{123, "Raymond", "Gitonga", "raytosh95@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(userId int64) (*User, *utils.AppError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.AppError) {
	log.Println("We are accessing the database")
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.AppError{
		Message:    fmt.Sprintf("User %v not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "Not found",
	}
}
