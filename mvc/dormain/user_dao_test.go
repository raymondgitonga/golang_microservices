package dormain

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	//Initialization

	//Execution
	user, err := GetUser(0)

	//validation
	assert.Nil(t, user, "We were not expecting user with id 0")
	assert.NotNil(t, err, "We were expecting an error with user id 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "Not found", err.Code)
}

func TestGetUserAUserIsFound(t *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(t, user, "We were not expecting user with id 123")
	assert.Nil(t, err, "We were not expecting an error wit user id 123")
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Raymond", user.FirstName)
}
