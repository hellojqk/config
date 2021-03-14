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

// ConfigStructController .
type ConfigStructController struct {
}

// Insert .
func (sc *ConfigStructController) Insert(c *gin.Context) {
	param := entity.ConfigStruct{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}

	result, err := service.ConfigStructInsertOne(context.Background(), param)
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
func (sc *ConfigStructController) UpdateOne(c *gin.Context) {
	var uriModel model.URIParam
	if err := c.ShouldBindUri(&uriModel); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}

	param := entity.ConfigStruct{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}

	result, err := service.ConfigStructUpdateOne(context.Background(), uriModel.StructKey, param)
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
func (sc *ConfigStructController) FindOne(c *gin.Context) {
	var uriModel model.URIParam
	if err := c.ShouldBindUri(&uriModel); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	fmt.Printf("key:%sc\n", uriModel.StructKey)
	result, err := service.ConfigStructFindOne(context.Background(), uriModel.StructKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	c.JSON(http.StatusOK, result)
}

// Find .
// /api/struct?page_num=1&page_size=10
func (sc *ConfigStructController) Find(c *gin.Context) {
	var param = entity.ListPagingParam{}
	if err := c.ShouldBind(&param); err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	util.PrintJSON("ConfigStruct Find", param)

	total, result, err := service.ConfigStructFind(context.Background(), param)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResult(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": total, "data": result})
}
