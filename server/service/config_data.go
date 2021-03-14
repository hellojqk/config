package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/hellojqk/config/server/entity"
	"github.com/hellojqk/config/server/repository"
	"github.com/hellojqk/config/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// getDataCollection .
func getDataCollection(ctx context.Context, structKey string) (config *entity.ConfigStruct, collection *mongo.Collection, err error) {
	config, err = ConfigStructFindOne(ctx, structKey)
	if err != nil {
		return nil, nil, err
	}
	if config == nil {
		return nil, nil, errors.New("config key not exists")
	}
	if !config.Array {
		collection = repository.DB.Collection("config_data")
	} else {
		collection = repository.DB.Collection(fmt.Sprintf("config_data_%s", structKey))
	}
	return
}

// ConfigDataInsertOne .
func ConfigDataInsertOne(ctx context.Context, structKey string, model entity.ConfigData) (result bool, err error) {

	config, collection, err := getDataCollection(ctx, structKey)
	if err != nil {
		return false, err
	}

	model.SetCreator("")
	//如果是结构配置的不是数组结构，强制model中的key为structKey
	if !config.Array {
		model["key"] = structKey
	}
	insertResult, err := collection.InsertOne(ctx, model)

	if err != nil {
		return false, err
	}

	return insertResult.InsertedID != nil, nil
}

// ConfigDataFindOne .
func ConfigDataFindOne(ctx context.Context, structKey string, dataKey string) (result interface{}, err error) {
	_, collection, err := getDataCollection(ctx, structKey)
	if err != nil {
		return nil, err
	}
	result = bson.M{}
	err = collection.FindOne(ctx, bson.M{"key": dataKey}).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ConfigDataFind .
func ConfigDataFind(ctx context.Context, structKey string, param entity.ListPagingParam) (total int64, result []entity.ConfigData, err error) {
	if param.Filter == nil {
		param.Filter = bson.M{}
	}
	if param.Sort == nil {
		param.Sort = bson.M{"creator.timestamp": -1}
	}

	util.PrintJSON("ConfigData Find", param)

	_, collection, err := getDataCollection(ctx, structKey)
	if err != nil {
		return 0, nil, err
	}

	total, err = collection.CountDocuments(ctx, param.Filter)
	if err != nil {
		return 0, nil, err
	}
	if total < 1 {
		return 0, nil, err
	}
	cur, err := collection.Find(ctx, param.Filter, options.Find().SetSort(param.Sort).SetLimit(param.PageSize).SetSkip((param.PageNum-1)*param.PageSize))

	if err != nil {
		return 0, nil, err
	}

	for cur.Next(ctx) {
		param := entity.ConfigData{}
		err := cur.Decode(&param)
		if err != nil {
			continue
		}
		result = append(result, param)
	}

	return total, result, nil
}

// ConfigDataUpdateOne .
func ConfigDataUpdateOne(ctx context.Context, key string, model entity.ConfigData) (result bool, err error) {
	_, collection, err := getDataCollection(ctx, key)
	if err != nil {
		return false, err
	}
	model.SetUpdater("")
	updateResult, err := collection.UpdateOne(ctx, bson.M{"key": key}, bson.M{"$set": model})
	if err != nil {
		return false, err
	}

	return updateResult.ModifiedCount == 1, nil
}
