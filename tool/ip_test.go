package tool

import (
	"fmt"
	"hh_tool/config"
	"hh_tool/model"
	"hh_tool/test"
	"reflect"
	"testing"
	"time"
)

func TestGetIpInfo(t *testing.T) {
	config.InitViper()
	tests := []struct {
		name string
	}{
		{"TestGetIpInfo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(GetIpInfo("65.49.130.105"))
		})
	}
}

func TestSaveIpInfoToMysql(t *testing.T) {
	test.InitTestConf()
	type args struct {
		ip      string
		ipiInfo string
		count   int64
	}
	tests := []struct {
		name string
		args args
	}{
		{"TestSaveIpInfoToMysql", args{"1123", "", 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveIpInfoToMysql(tt.args.ip, tt.args.ipiInfo, tt.args.count)
		})
	}
}

func TestBuildSaveIpInfo(t *testing.T) {
	type args struct {
		ip     string
		ipInfo string
		count  int64
	}
	tests := []struct {
		name string
		args args
		want model.V2rayAccessLogIpCount
	}{
		{name: "TestBuildIpInfo", args: args{ip: "114.245.147.44", count: 2}, want: model.V2rayAccessLogIpCount{Ip: "114.245.147.44", Count: 2, UpdateTime: time.Now().Truncate(time.Second)}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildSaveIpInfo(tt.args.ip, tt.args.ipInfo, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildSaveIpInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
