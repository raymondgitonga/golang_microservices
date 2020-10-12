package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/raymondgitonga/golang_microservices/mvc/service"
	"github.com/raymondgitonga/golang_microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)

	if err != nil {
		apiErr := &utils.AppError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := service.UserService.GetUser(userId)

	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)
}
