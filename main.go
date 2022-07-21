package main

import (
	"flag"
	"hh_tool/config"
	"hh_tool/database"
	"hh_tool/model"
	"hh_tool/test"
	d "hh_tool/web3/domain"
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
	var domain int
	flag.IntVar(&aid, "aid", 0, "access_log id")
	flag.IntVar(&domain, "bd", 0, "run bit domian")
	flag.Parse()
	if aid > 0 {
		test.RerunAccessLogHistoryRedisData(aid)
		os.Exit(0)
	}
	if domain > 0 {
		d.GetDomainInfo()
		os.Exit(0)
	}
}
