package service

import (
	"context"
	"testing"

	"github.com/hellojqk/config/entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestConfigDataInsertOne(t *testing.T) {
	result, err := ConfigDataInsertOne(context.Background(), "key2", entity.ConfigData{"key": "key2", "qty": 100})
	assert.Equal(t, nil, err)
	t.Logf("%v\n", result)
}

func TestConfigDataFindOne(t *testing.T) {
	result, err := ConfigDataFindOne(context.Background(), "key1", "key1")
	assert.Equal(t, nil, err)
	t.Logf("%v\n", result)
}

func TestConfigDataFind(t *testing.T) {
	total, result, err := ConfigDataFind(context.Background(), "form", entity.ListPagingParam{PageNum: 1, PageSize: 10, Filter: bson.M{}, Sort: bson.M{}})
	assert.Equal(t, nil, err)
	t.Logf("%d\n%v\n", total, result)
}
