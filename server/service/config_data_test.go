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
