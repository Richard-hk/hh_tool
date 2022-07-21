package util

import "os"

const (
	V2RAY = "v2ray"
)

var MonitorApp = []string{V2RAY}
var StopSignal = make(chan os.Signal, 1)

var BitDomain_UnKnown = -1 // 未知
// var BitDomain_Register = 1  // 已经注册
// var BitDomain_Sale = 2      // 出售
// var BitDomain_Available = 3 // 可用
// var BitDomain_Reserved = 4  // 保护中
