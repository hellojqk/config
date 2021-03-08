package service

import (
	"context"

	"github.com/hellojqk/config/entity"
	"github.com/hellojqk/config/repository"
	util "github.com/hellojqk/config/tools/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConfigStruct .
type ConfigStruct struct {
}

func newCollection() *mongo.Collection {
	return repository.DB.Collection("struct_config")
}

var structConfigService = ConfigStruct{}

// InsertOne .
func (s *ConfigStruct) InsertOne(ctx context.Context, model entity.ConfigStruct) (result interface{}, err error) {
	model.Create(0)
	insertResult, err := newCollection().InsertOne(ctx, model)
	if err != nil {
		return nil, err
	}

	return insertResult.InsertedID, nil
}

// FindOne .
func (s *ConfigStruct) FindOne(ctx context.Context, structKey string) (result *entity.ConfigStruct, err error) {
	result = &entity.ConfigStruct{}
	err = newCollection().FindOne(ctx, bson.M{"key": structKey}).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Find .
func (s *ConfigStruct) Find(ctx context.Context, param entity.ListPagingParam) (result []entity.ConfigStruct, err error) {
	if param.Filter == nil {
		param.Filter = bson.M{}
	}
	if param.Sort == nil {
		param.Sort = bson.M{}
	}
	util.PrintJSON("ConfigStruct Find", param)
	cur, err := newCollection().Find(ctx, param.Filter, options.Find().SetSort(param.Sort).SetLimit(param.PageSize).SetSkip((param.PageNum-1)*param.PageSize))

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		model := entity.ConfigStruct{}
		err := cur.Decode(&model)
		if err != nil {
			continue
		}
		result = append(result, model)
	}

	return result, nil
}

// UpdateOne .
func (s *ConfigStruct) UpdateOne(ctx context.Context, structKey string, model entity.ConfigStruct) (result interface{}, err error) {
	model.Update(0)
	updateResult, err := newCollection().UpdateOne(ctx, bson.M{"key": structKey}, bson.M{"$set": bson.M{
		"title":          model.Title,
		"description":    model.Description,
		"secret":         model.Secret,
		"array":          model.Array,
		"schema":         model.Schema,
		"update_time":    model.UpdateTime,
		"update_user_id": model.UpdateUserID,
	}})
	if err != nil {
		return nil, err
	}

	return updateResult.ModifiedCount, nil
}
