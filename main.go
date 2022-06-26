package main

import (
	"hh_tool/config"
	"hh_tool/database"
	"hh_tool/model"
)

func init() {
	config.InitViper()
	config.InitLogrus()
	model.InitMysqlCon()
	database.InitRedisCon()
}

func main() {
	Run()
	select {}
}
