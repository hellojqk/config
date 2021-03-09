package service

import (
	"context"
	"testing"

	"github.com/hellojqk/config/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestConfigData_InsertOne(t *testing.T) {
	result, err := structDataService.InsertOne(context.Background(), "key1", entity.ConfigData{Key: "key1", Data: bson.M{"key": "key1", "qty": 100}})
	assert.Equal(t, nil, err)
	t.Logf("%v\n", result)
}

func TestConfigData_FindOne(t *testing.T) {
	result, err := structDataService.FindOne(context.Background(), "key1", "key1")
	assert.Equal(t, nil, err)
	t.Logf("%v\n", result)
}

func TestConfigData_Find(t *testing.T) {
	total, result, err := structDataService.Find(context.Background(), "aa", entity.ListPagingParam{PageNum: 1, PageSize: 10, Filter: bson.M{}, Sort: bson.M{}})
	assert.Equal(t, nil, err)
	t.Logf("%d\n%v\n", total, result)
}
