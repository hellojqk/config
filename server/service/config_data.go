package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/hellojqk/config/entity"
	"github.com/hellojqk/config/repository"
	util "github.com/hellojqk/config/tools/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConfigData .
type ConfigData struct {
}

var structDataService = ConfigData{}

// getDataCollection .
func getDataCollection(ctx context.Context, structKey string) (config *entity.ConfigStruct, collection *mongo.Collection, err error) {
	config, err = structConfigService.FindOne(ctx, structKey)
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

// InsertOne .
func (s *ConfigData) InsertOne(ctx context.Context, structKey string, model entity.ConfigData) (result interface{}, err error) {

	_, collection, err := getDataCollection(ctx, structKey)
	if err != nil {
		return nil, err
	}

	model.Create(0)
	insertResult, err := collection.InsertOne(ctx, bson.M{"key": model.Key, "data": model.Data})

	if err != nil {
		return nil, err
	}

	return insertResult.InsertedID, nil
}

// FindOne .
func (s *ConfigData) FindOne(ctx context.Context, structKey string, dataKey string) (result interface{}, err error) {
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

// Find .
func (s *ConfigData) Find(ctx context.Context, structKey string, param entity.ListPagingParam) (total int64, result []entity.ConfigData, err error) {
	if param.Filter == nil {
		param.Filter = bson.M{}
	}
	if param.Sort == nil {
		param.Sort = bson.M{}
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
		model := entity.ConfigData{}
		err := cur.Decode(&model)
		if err != nil {
			continue
		}
		result = append(result, model)
	}

	return total, result, nil
}

// UpdateOne .
func (s *ConfigData) UpdateOne(ctx context.Context, key string, model entity.ConfigData) (result interface{}, err error) {
	model.Update(0)
	_, collection, err := getDataCollection(ctx, key)
	if err != nil {
		return nil, err
	}
	updateResult, err := collection.UpdateOne(ctx, bson.M{"key": key}, bson.M{"$set": bson.M{
		"data":           model.Data,
		"update_time":    model.UpdateTime,
		"update_user_id": model.UpdateUserID,
	}})
	if err != nil {
		return nil, err
	}

	return updateResult.ModifiedCount, nil
}
