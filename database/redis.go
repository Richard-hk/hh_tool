package database

import (
	"hh_tool/util"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitRedisCon() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
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
