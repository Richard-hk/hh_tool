package test

import (
	"fmt"
	"hh_tool/model"
	"hh_tool/tool"
	"sync"
	"time"
)

func RerunAccessLogHistoryRedisData(MaxId int) {
	var v2ray_access_log model.V2rayAccessLog
	var wg *sync.WaitGroup
	startTime := time.Now().Unix()
	for i := 1; i <= MaxId; i = i + 100 {
		wg.Add(1)
		go func(i int) {
			v2rayAccessLogs, _ := v2ray_access_log.GetV2rayAccessLogById(int64(i), int64(i+99))
			for _, v := range v2rayAccessLogs {
				fmt.Println(v.Ip)
				tool.SaveIpInfo(v.Ip)
				wg.Done()
			}
		}(i)
	}
	wg.Done()
	endTime := time.Now().Unix()
	fmt.Println("RerunAccessLogHistoryRedisData finished, total spend time", endTime-startTime, " s")
}
