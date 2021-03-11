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

// ConfigData .
type ConfigData struct {
}

var structDataService = service.ConfigData{}

// Insert .
func (s *ConfigData) Insert(c *gin.Context) {

	var uriModel model.URI
	if err := c.ShouldBindUri(&uriModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	model := entity.ConfigData{}
	if err := c.ShouldBind(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	result, err := structDataService.InsertOne(context.Background(), uriModel.StructKey, model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"err": "保存失败"})
	}
	c.Status(http.StatusOK)
}

// UpdateOne .
func (s *ConfigData) UpdateOne(c *gin.Context) {
	var uriModel model.URI
	if err := c.ShouldBindUri(&uriModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	model := entity.ConfigData{}
	if err := c.ShouldBind(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	result, err := structDataService.UpdateOne(context.Background(), uriModel.StructKey, model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"err": "保存失败"})
	}
	c.Status(http.StatusOK)
}

// FindOne .
func (s *ConfigData) FindOne(c *gin.Context) {
	var uriModel model.URI
	if err := c.ShouldBindUri(&uriModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	fmt.Printf("key:%s\t%s\n", uriModel.StructKey, uriModel.DataKey)
	result, err := structDataService.FindOne(context.Background(), uriModel.StructKey, uriModel.DataKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Find .
// /api/struct?page_num=1&page_size=10
func (s *ConfigData) Find(c *gin.Context) {
	var uriModel model.URI
	if err := c.ShouldBindUri(&uriModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var param = entity.ListPagingParam{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	util.PrintJSON("ConfigData Find", param)

	total, result, err := structDataService.Find(context.Background(), uriModel.StructKey, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": total, "data": result})
}
