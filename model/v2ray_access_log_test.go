package model

import (
	"fmt"
	"hh_tool/config"
	"testing"
)

func TestV2rayAccessLog_GetV2rayAccessLogById(t *testing.T) {
	config.InitLogrus()
	config.InitViper()
	config.InitRedisCon()
	config.InitMysqlCon()
	var v2rayAccessLog V2rayAccessLog
	got, err := v2rayAccessLog.GetV2rayAccessLogById(1, 2)
	fmt.Println(got, err)
}
