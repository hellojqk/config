package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hellojqk/config/repository"
	"github.com/hellojqk/config/server/controller"
	"github.com/spf13/viper"
)

// Run .
func Run() {
	repository.InitConn()
	g := gin.Default()
	apiGroup := g.Group("api")

	structConfigController := &controller.ConfigStruct{}
	apiGroup.POST("/struct", structConfigController.Insert)
	apiGroup.PUT("/struct/:struct_key", structConfigController.UpdateOne)
	apiGroup.GET("/struct/:struct_key", structConfigController.FindOne)
	apiGroup.GET("/struct", structConfigController.Find)

	structDataController := &controller.ConfigData{}
	apiGroup.POST("/struct/:struct_key/data", structDataController.Insert)
	apiGroup.PUT("/struct/:struct_key/data/:data_key", structDataController.UpdateOne)
	apiGroup.GET("/struct/:struct_key/data/:data_key", structDataController.FindOne)
	apiGroup.GET("/struct/:struct_key/data", structDataController.Find)

	g.Run(fmt.Sprintf(":%d", viper.GetInt("serverPort")))
}
