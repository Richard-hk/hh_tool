package test

import (
	"fmt"
	"hh_tool/model"
	"hh_tool/monitor"
	"hh_tool/tool"
	"time"
)

func RerunAccessLogHistoryRedisData(MaxId int) {
	var v2ray_access_log model.V2rayAccessLog
	startTime := time.Now().Unix()
	v2rayIpCountMap := make(map[string]int)
	for i := 1; i <= MaxId; i = i + 100 {
		fmt.Println("i: ", i)
		v2rayAccessLogs, _ := v2ray_access_log.GetV2rayAccessLogById(int64(i), int64(i+99))
		for _, v := range v2rayAccessLogs {
			monitor.BuildV2rayIpCountMap(v2rayIpCountMap, v)
		}
		tool.SaveIpInfo(v2rayIpCountMap)
		v2rayIpCountMap = nil
	}
	endTime := time.Now().Unix()
	fmt.Println("RerunAccessLogHistoryRedisData finished, total spend time", endTime-startTime, " s")
}
