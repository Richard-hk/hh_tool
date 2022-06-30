package test

import (
	"fmt"
	"hh_tool/model"
	"hh_tool/tool"
	"sync"
	"time"

	"github.com/spf13/viper"
)

func RerunAccessLogHistoryRedisData(MaxId int) {
	var v2ray_access_log model.V2rayAccessLog
	var wg sync.WaitGroup
	startTime := time.Now().Unix()
	maxNum := make(chan struct{}, viper.GetInt("num.connect_mysql"))
	for i := 1; i <= MaxId; i = i + 100 {
		wg.Add(1)
		maxNum <- struct{}{}
		fmt.Println("i: ", i)
		go func(i int) {
			v2rayAccessLogs, _ := v2ray_access_log.GetV2rayAccessLogById(int64(i), int64(i+99))
			for _, v := range v2rayAccessLogs {
				tool.SaveIpInfo(v.Ip)
			}
			wg.Done()
			<-maxNum
			fmt.Println("go runtine finished i:", i)
		}(i)
	}
	wg.Wait()
	endTime := time.Now().Unix()
	fmt.Println("RerunAccessLogHistoryRedisData finished, total spend time", endTime-startTime, " s")
}
