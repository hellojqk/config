package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/hellojqk/config/util"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var redisCli *redis.Client

func init() {
	util.WaitInitFuncsAdd(initRedisClient)
}

func initRedisClient() error {
	// var err error
	redisCli = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Username: viper.GetString("redis.userName"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})
	_, err := redisCli.Ping(context.Background()).Result()
	if err != nil {
		return errors.WithMessage(err, "redis client init error")
	}
	return nil
}
