package test

import (
	"fmt"
	"hh_tool/config"
	"testing"
)

func TestSetRedisVal(t *testing.T) {
	rdb := config.InitRedisCon()
	val2, err := rdb.HGet("ip", "123").Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val2)
}
