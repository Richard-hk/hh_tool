package v2ray

import (
	"hh_tool/config"
	"testing"
)

func TestMonitorV2rayAccessLogIpCount(t *testing.T) {
	config.Init()
	MonitorV2rayAccessLogIpCount()
}
