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
	for i := 1; i <= MaxId; i = i + 100 {
		v2rayAccessLogs, _ := v2ray_access_log.GetV2rayAccessLogById(int64(i), int64(i+99))
		v2rayIpCountMap := make(map[string]int)
		for _, v := range v2rayAccessLogs {
			monitor.BuildV2rayIpCountMap(v2rayIpCountMap, v)
		}
		tool.SaveIpInfo(v2rayIpCountMap)
		fmt.Println("go runtine finished i:", i)
	}
	endTime := time.Now().Unix()
	fmt.Println("RerunAccessLogHistoryRedisData finished, total spend time", endTime-startTime, " s")
}

// var v2ray_access_log model.V2rayAccessLog
// var wg sync.WaitGroup
// startTime := time.Now().Unix()
// var mutex sync.Mutex
// maxNum := make(chan struct{}, viper.GetInt("num.connect_mysql"))
// for i := 1; i <= MaxId; i = i + 100 {
// 	wg.Add(1)
// 	maxNum <- struct{}{}
// 	go func(i int) {
// 		v2rayAccessLogs, _ := v2ray_access_log.GetV2rayAccessLogById(int64(i), int64(i+99))
// 		v2rayIpCountMap := make(map[string]int)
// 		for _, v := range v2rayAccessLogs {
// 			monitor.BuildV2rayIpCountMap(v2rayIpCountMap, v)
// 		}
// 		tool.SaveIpInfo(v2rayIpCountMap)
// 		wg.Done()
// 		<-maxNum
// 		fmt.Println("go runtine finished i:", i)
// 	}(i)
// }
// wg.Wait()
// endTime := time.Now().Unix()
