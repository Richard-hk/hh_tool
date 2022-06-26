package test

import (
	"hh_tool/config"
	"hh_tool/model"
)

func InitTestConf() {
	config.InitLogrus()
	config.InitViper()
	model.InitMysqlCon()
}
