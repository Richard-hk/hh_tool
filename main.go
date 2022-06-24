package main

import (
	"hh_tool/config"
	"hh_tool/model"
)

func init() {
	config.InitViper()
	config.InitLogrus()
	model.InitMysqlCon()
}

func main() {
	Run()
	select {}
}
