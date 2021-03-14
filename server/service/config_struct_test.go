package service

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hellojqk/config/entity"
	"github.com/hellojqk/config/util"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".hpa" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".config")
	viper.AddConfigPath("config/")
	viper.AddConfigPath("../config/")
	viper.AddConfigPath("../../config/")
	viper.AddConfigPath("../../../config/")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func TestMain(m *testing.M) {
	initConfig()
	util.WaitInitFuncsExec()
	os.Exit(m.Run())
}

func TestConfigStructInsertOne(t *testing.T) {
	result, err := ConfigStructInsertOne(context.Background(), entity.ConfigStruct{Key: "key11"})
	result, err = ConfigStructInsertOne(context.Background(), entity.ConfigStruct{Key: "key2"})
	result, err = ConfigStructInsertOne(context.Background(), entity.ConfigStruct{Key: "key3"})
	result, err = ConfigStructInsertOne(context.Background(), entity.ConfigStruct{Key: "key4"})
	result, err = ConfigStructInsertOne(context.Background(), entity.ConfigStruct{Key: "key5"})
	result, err = ConfigStructInsertOne(context.Background(), entity.ConfigStruct{Key: "key6"})
	assert.Equal(t, nil, err)
	t.Logf("%v\n", result)
}

func TestConfigStructFindOne(t *testing.T) {
	result, err := ConfigStructFindOne(context.Background(), "key1")
	assert.Equal(t, nil, err)
	t.Logf("%v\n", result)
}

func TestConfigStructFind(t *testing.T) {
	total, result, err := ConfigStructFind(context.Background(), entity.ListPagingParam{PageNum: 1, PageSize: 10, Filter: bson.M{"key": bson.M{"$regex": "key"}}})
	assert.Equal(t, nil, err)
	t.Logf("%d\n%v\n", total, result)
}

func TestConfigStructUpdateOne(t *testing.T) {
	result, err := ConfigStructUpdateOne(context.Background(), "key11", entity.ConfigStruct{Title: "title"})
	assert.Equal(t, nil, err)
	t.Logf("%v\n", result)
}

func TestConfigStructEdit(t *testing.T) {
	result, err := ConfigStructInsertOne(context.Background(), entity.ConfigStruct{Key: "key111"})

	result, err = ConfigStructUpdateOne(context.Background(), "key111", entity.ConfigStruct{Title: "title"})
	assert.Equal(t, nil, err)
	t.Logf("%v\n", result)
}
