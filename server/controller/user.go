package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellojqk/config/server/model"
	"github.com/hellojqk/config/server/service"
)

// UserController .
type UserController struct {
}

// Login .
func (uc *UserController) Login(c *gin.Context) {
	param := model.UserLoginParam{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}

	jwtToken, err := service.UserLoginParam(context.Background(), param)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	c.Header("Authorization", fmt.Sprintf("Bearer %s", jwtToken))
	c.JSON(http.StatusOK, gin.H{"token": jwtToken})
}

// InsertOne .
func (uc *UserController) InsertOne(c *gin.Context) {
	return
}

func (uc *UserController) JWT(c *gin.Context) {

	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		c.JSON(http.StatusBadRequest, gin.H{"err": "lost Authorization Token"})
		c.Abort()
		return
	}
	userKey, err := service.UserTokenValid(context.Background(), authorization)

	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		c.Abort()
		return
	}
	c.Set("userKey", userKey)
	c.Next()
}
