package service

import (
	"github.com/raymondgitonga/golang_microservices/mvc/dormain"
	"github.com/raymondgitonga/golang_microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type userDaoMock struct{}

var (
	UserDaoMock     userDaoMock
	getUserFunction func(userId int64) (*dormain.User, *utils.AppError)
)

func init() {
	dormain.UserDao = &userDaoMock{}
}

func (u *userDaoMock) GetUser(userId int64) (*dormain.User, *utils.AppError) {
	return getUserFunction(userId)
}
func TestGetUserNotInDataBase(t *testing.T) {
	getUserFunction = func(userId int64) (*dormain.User, *utils.AppError) {
		return nil, &utils.AppError{
			StatusCode: http.StatusNotFound,
		}
	}
	user, err := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUserPresentInDataBase(t *testing.T) {
	getUserFunction = func(userId int64) (*dormain.User, *utils.AppError) {
		return &dormain.User{
			Id: 123,
		}, nil
	}
	user, err := UserService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, int64(123), user.Id)
}
