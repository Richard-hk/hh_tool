package main

import (
	"flag"
	"hh_tool/config"
	"hh_tool/database"
	"hh_tool/model"
	"hh_tool/test"
	"os"
)

func init() {
	config.InitViper()
	config.InitLogrus()
	model.InitMysqlCon()
	database.InitRedisCon()
}

func main() {
	manual()
	Run()
	select {}
}

func manual() {
	var aid int
	flag.IntVar(&aid, "aid", 1, "access_log id")
	flag.Parse()
	if aid > 0 {
		test.RerunAccessLogHistoryRedisData(aid)
		os.Exit(0)
	}
}
