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

	var m model.URI
	if err := c.ShouldBindUri(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	model := entity.ConfigData{}
	if err := c.ShouldBind(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	result, err := structDataService.InsertOne(context.Background(), m.StructKey, model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// UpdateOne .
func (s *ConfigData) UpdateOne(c *gin.Context) {
	var m model.URI
	if err := c.ShouldBindUri(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	model := entity.ConfigData{}
	if err := c.ShouldBind(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	result, err := structDataService.UpdateOne(context.Background(), m.StructKey, model)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// FindOne .
func (s *ConfigData) FindOne(c *gin.Context) {
	var m model.URI
	if err := c.ShouldBindUri(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	fmt.Printf("key:%s\t%s\n", m.StructKey, m.DataKey)
	result, err := structDataService.FindOne(context.Background(), m.StructKey, m.DataKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Find .
// /api/struct?page_num=1&page_size=10
func (s *ConfigData) Find(c *gin.Context) {
	var param = entity.ListPagingParam{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	util.PrintJSON("ConfigData Find", param)

	result, err := structDataService.Find(context.Background(), param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
