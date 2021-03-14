package service

import (
	"context"

	"github.com/hellojqk/config/server/entity"
	"github.com/hellojqk/config/server/repository"
	"github.com/hellojqk/config/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConfigStruct .
type ConfigStruct struct {
}

func newCollection() *mongo.Collection {
	return repository.DB.Collection("config_struct")
}

// ConfigStructInsertOne .
func ConfigStructInsertOne(ctx context.Context, model entity.ConfigStruct) (result bool, err error) {
	model.SetCreator("")
	insertResult, err := newCollection().InsertOne(ctx, model)
	if err != nil {
		return false, err
	}

	return insertResult.InsertedID != nil, nil
}

// ConfigStructFindOne .
func ConfigStructFindOne(ctx context.Context, structKey string) (result *entity.ConfigStruct, err error) {
	result = &entity.ConfigStruct{}
	err = newCollection().FindOne(ctx, bson.M{"key": structKey}).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ConfigStructFind .
func ConfigStructFind(ctx context.Context, param entity.ListPagingParam) (total int64, result []entity.ConfigStruct, err error) {
	if param.Filter == nil {
		param.Filter = bson.M{}
	}
	if param.Sort == nil {
		param.Sort = bson.M{"creator.timestamp": -1}
	}
	util.PrintJSON("ConfigStruct Find", param)
	total, err = newCollection().CountDocuments(ctx, param.Filter)
	if err != nil {
		return 0, nil, err
	}
	if total < 1 {
		return 0, nil, err
	}
	cur, err := newCollection().Find(ctx, param.Filter, options.Find().SetSort(param.Sort).SetLimit(param.PageSize).SetSkip((param.PageNum-1)*param.PageSize))

	if err != nil {
		return 0, nil, err
	}

	for cur.Next(ctx) {
		param := entity.ConfigStruct{}
		err := cur.Decode(&param)
		if err != nil {
			continue
		}
		result = append(result, param)
	}

	return total, result, nil
}

// UpdateOne .
func ConfigStructUpdateOne(ctx context.Context, structKey string, model entity.ConfigStruct) (result bool, err error) {
	model.SetUpdater("")
	updateResult, err := newCollection().UpdateOne(ctx, bson.M{"key": structKey}, bson.M{"$set": bson.M{
		"title":       model.Title,
		"description": model.Description,
		"secret":      model.Secret,
		"array":       model.Array,
		"schema":      model.Schema,
		"updater":     model.Updater,
	}})
	if err != nil {
		return false, err
	}

	return updateResult.ModifiedCount == 1, nil
}
