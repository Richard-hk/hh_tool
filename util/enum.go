package util

import "os"

const (
	V2RAY     = "v2ray"
	BitDomain = "bitDomain"
)

var MonitorApp = []string{V2RAY}
var MonitorWeb = []string{BitDomain}
var StopSignal = make(chan os.Signal, 1)

var BitDomain_UnKnown = -2     // 未知
var BitDomain_UnSale = -1      // 未开放注册
var BitDomain_Available_0 = 0  // 可用
var BitDomain_Available_1 = 1  // 可用
var BitDomain_Register_6 = 6   // 已经注册
var BitDomain_Reserved = 7     // 保护中
var BitDomain_Sale = 8         // 出售
var BitDomain_UnExist = 14     // 未锻造
var BitDomain_Register_15 = 15 // 已经注册

var BitDomainMap = map[int]string{
	BitDomain_UnKnown:     "未知",
	BitDomain_UnSale:      "未释放",
	BitDomain_Available_0: "已经注册",
	BitDomain_Register_6:  "已经注册",
	BitDomain_Sale:        "出售",
	BitDomain_Available_0: "可用",
	BitDomain_Available_1: "可用",
	BitDomain_Reserved:    "保护中",
	BitDomain_UnExist:     "未锻造",
}
