package common

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var Redis *redis.Client

func init() {
	addr := viper.GetString("redis.addr")
	pwd := viper.GetString("redis.pwd")

	Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
	})
}
