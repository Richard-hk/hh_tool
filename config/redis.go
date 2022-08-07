package config

import (
	"hh_tool/util"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func InitRedisCon() {
	rdb = SetRedisCon()
}

func SetRedisCon() (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("connect.redis.adr"),
		Password: viper.GetString("connect.redis.pwd"),
		DB:       viper.GetInt("connect.redis.db"),
	})

	_, err := rdb.Ping().Result()
	util.HandleError(err, "InitRedis failed")
	return rdb
}

func GetRedisCon() *redis.Client {
	if rdb == nil {
		panic("GetRedisCon failed")
	}
	return rdb
}
