package model

import (
	"hh_tool/database"

	"gorm.io/gorm"
)

var hhToolDb *gorm.DB

func InitMysqlCon() {
	setHhToolCon()
}

func setHhToolCon() {
	hhToolDb = database.SetMysqlCon("hh_tool")
}

func GetHhToolCon() *gorm.DB {
	if hhToolDb == nil {
		panic("connect hh_tool failed")
	}
	return hhToolDb
}
