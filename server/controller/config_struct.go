package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellojqk/config/entity"
	"github.com/hellojqk/config/server/model"
	"github.com/hellojqk/config/server/service"
	util "github.com/hellojqk/config/tools/utils"
)

// ConfigStruct .
type ConfigStruct struct {
}

var structConfigService = service.ConfigStruct{}

// Insert .
func (s *ConfigStruct) Insert(c *gin.Context) {
	model := entity.ConfigStruct{}
	if err := c.ShouldBind(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	result, err := structConfigService.InsertOne(context.Background(), model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// UpdateOne .
func (s *ConfigStruct) UpdateOne(c *gin.Context) {
	var m model.URI
	if err := c.ShouldBindUri(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	model := entity.ConfigStruct{}
	if err := c.ShouldBind(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	result, err := structConfigService.UpdateOne(context.Background(), m.StructKey, model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// FindOne .
func (s *ConfigStruct) FindOne(c *gin.Context) {
	var m model.URI
	if err := c.ShouldBindUri(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	fmt.Printf("key:%s\n", m.StructKey)
	result, err := structConfigService.FindOne(context.Background(), m.StructKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Find .
// /api/struct?page_num=1&page_size=10
func (s *ConfigStruct) Find(c *gin.Context) {
	var param = entity.ListPagingParam{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	util.PrintJSON("ConfigStruct Find", param)

	result, err := structConfigService.Find(context.Background(), param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}