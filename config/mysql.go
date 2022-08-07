package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var hhToolDb *gorm.DB

func InitMysqlCon() {
	setHhToolCon()
}

func SetMysqlCon(dbName string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("connect.mysql.usr"),
		viper.GetString("connect.mysql.pwd"),
		viper.GetString("connect.mysql.host"),
		viper.GetInt("connect.mysql.port"),
		dbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error("connect to mysql error dbname is ", dbName)
	}
	return db
}

func setHhToolCon() {
	hhToolDb = SetMysqlCon("hh_tool")
}

func GetHhToolCon() *gorm.DB {
	if hhToolDb == nil {
		panic("connect hh_tool failed")
	}
	return hhToolDb
}
