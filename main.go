package main

import (
	"flag"
	"hh_tool/config"
	"hh_tool/test"
	"os"
)

func init() {
	config.Init()
}

func main() {
	manual()
	Run()
	select {}
}

func manual() {
	var aid, domain, telegram int
	flag.IntVar(&aid, "aid", 0, "access_log id")
	flag.IntVar(&domain, "bd", 0, "run bit domian")
	flag.IntVar(&telegram, "tg", 0, "run telegram bot")
	flag.Parse()
	if aid > 0 {
		test.RerunAccessLogHistoryRedisData(aid)
		os.Exit(0)
	}
}
