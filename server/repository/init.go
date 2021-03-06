package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/hellojqk/config/util"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

// CLI .
var CLI *mongo.Client

// DB .
var DB *mongo.Database

func init() {
	util.WaitInitFuncsAdd(initMongoClient)
}

// initMongoClient 初始化连接
func initMongoClient() (err error) {
	connectionStr := viper.GetString("connectionString")
	if connectionStr == "" {
		return errors.New("connectionStr is null")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	CLI, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionStr))
	if err != nil {
		return errors.WithMessage(err, "mongodb Connect error")
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = CLI.Ping(ctx, readpref.Primary())
	if err != nil {
		return errors.WithMessage(err, "mongodb Ping error")
	}

	var dataBaseName = viper.GetString("dataBase")
	DB = CLI.Database(dataBaseName)

	collectionNames, err := DB.ListCollectionNames(context.Background(), bsonx.Doc{})

	var initCollections = []string{"config_struct", "config_data", "role", "user", "group"}
	option := options.CreateCollection()
	fmt.Printf("%v\n", collectionNames)
	for _, collection := range initCollections {
		if !util.ExistsInStringArray(collection, collectionNames) {
			err := DB.CreateCollection(ctx, collection, option)
			if err != nil {
				return errors.WithMessage(err, "mongodb Collection init")
			}
		}
	}

	for _, collection := range initCollections {
		_, err := DB.Collection(collection).Indexes().CreateOne(context.Background(), mongo.IndexModel{
			Keys:    bsonx.Doc{{Key: "key", Value: bsonx.Int32(1)}},
			Options: options.Index().SetUnique(true),
		})
		if err != nil {
			return errors.WithMessage(err, "mongodb index init")
		}
	}
	return nil
}

// Close .
func Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err := CLI.Disconnect(ctx)
	if err != nil {
		panic(errors.Wrap(err, "mongodb Disconnect error"))
	}
}
