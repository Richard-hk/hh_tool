package v2ray

import (
	"fmt"
	"hh_tool/model"
	"hh_tool/util"
)

func MonitorV2rayAccessLogIpCount() {
	dt := util.GetNowDt()
	v2rayAccessLogIpCounts, _ := new(model.V2rayAccessLog).GetV2rayAccessLogIpCountByDt(dt, util.V2rayAccessLogStatusAccepted)
	fmt.Println(v2rayAccessLogIpCounts)
}
