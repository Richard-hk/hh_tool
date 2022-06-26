package test

import (
	"hh_tool/config"
	"hh_tool/database"
	"hh_tool/model"
)

func InitTestConf() {
	config.InitLogrus()
	config.InitViper()
	model.InitMysqlCon()
	database.InitRedisCon()
}
