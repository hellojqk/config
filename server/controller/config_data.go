package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellojqk/config/server/entity"
	"github.com/hellojqk/config/server/model"
	"github.com/hellojqk/config/server/service"
	"github.com/hellojqk/config/util"
)

// ConfigDataController .
type ConfigDataController struct {
}

// Insert .
func (dc *ConfigDataController) Insert(c *gin.Context) {

	var uriModel model.URIParam
	if err := c.ShouldBindUri(&uriModel); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}

	param := entity.ConfigData{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}

	result, err := service.ConfigDataInsertOne(context.Background(), uriModel.StructKey, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"err": "保存失败"})
	}
	c.Status(http.StatusOK)
}

// UpdateOne .
func (dc *ConfigDataController) UpdateOne(c *gin.Context) {
	var uriModel model.URIParam
	if err := c.ShouldBindUri(&uriModel); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}

	param := entity.ConfigData{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}

	result, err := service.ConfigDataUpdateOne(context.Background(), uriModel.StructKey, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	if !result {
		c.JSON(http.StatusBadRequest, gin.H{"err": "保存失败"})
	}
	c.Status(http.StatusOK)
}

// FindOne .
func (dc *ConfigDataController) FindOne(c *gin.Context) {
	var uriModel model.URIParam
	if err := c.ShouldBindUri(&uriModel); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	fmt.Printf("key:%dc\t%dc\n", uriModel.StructKey, uriModel.DataKey)
	result, err := service.ConfigDataFindOne(context.Background(), uriModel.StructKey, uriModel.DataKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// Find .
// /api/struct?page_num=1&page_size=10
func (dc *ConfigDataController) Find(c *gin.Context) {
	var uriModel model.URIParam
	if err := c.ShouldBindUri(&uriModel); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}

	var param = entity.ListPagingParam{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	util.PrintJSON("ConfigData Find", param)

	total, result, err := service.ConfigDataFind(context.Background(), uriModel.StructKey, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": total, "data": result})
}
