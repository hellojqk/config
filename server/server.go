package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hellojqk/config/server/controller"
	"github.com/spf13/viper"
)

// Run .
func Run() {
	g := gin.Default()

	userController := controller.UserController{}

	apiGroup := g.Group("api")
	apiGroup.GET("/user/login", userController.Login)

	apiStructGroup := apiGroup.Group("struct", userController.JWT)
	structConfigController := &controller.ConfigStructController{}
	apiStructGroup.POST("/", structConfigController.Insert)
	apiStructGroup.PUT("/:struct_key", structConfigController.UpdateOne)
	apiStructGroup.GET("/:struct_key", structConfigController.FindOne)
	apiStructGroup.GET("/", structConfigController.Find)

	structDataController := &controller.ConfigDataController{}
	apiStructGroup.POST("/:struct_key/data", structDataController.Insert)
	apiStructGroup.PUT("/:struct_key/data/:data_key", structDataController.UpdateOne)
	apiStructGroup.GET("/:struct_key/data/:data_key", structDataController.FindOne)
	apiStructGroup.GET("/:struct_key/data", structDataController.Find)

	openAPIGroup := g.Group("openapi")
	openAPIGroup.GET("/config/:struct_key")
	openAPIGroup.GET("/config/:struct_key/:data_key")

	g.Run(fmt.Sprintf(":%d", viper.GetInt("serverPort")))
}
