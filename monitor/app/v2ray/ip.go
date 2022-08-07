package v2ray

import (
	"hh_tool/config"
	"hh_tool/model"
	"hh_tool/tool/urlinfo"
	"hh_tool/util"
	"time"

	"github.com/spf13/viper"
)

func SaveIpInfo(v2rayIpCountMap map[string]int) {
	for ip, valCount := range v2rayIpCountMap {
		redisInfoCount, err := GetRedisIpCount(ip)
		var count = int64(valCount)
		if err == nil {
			count += redisInfoCount
		}
		SetRedisIpCount(ip, count)
		if err == nil {
			UpdateIpInfoToMysql(ip, count)
			continue
		}
		ipInfo := GetIpInfo(ip)
		SaveIpInfoToMysql(ip, ipInfo, int64(count))
	}
}

func GetIpInfo(ip string) string {
	ipSite := viper.GetString("url.ip.url")
	ipUrl := ipSite + "/" + ip + ".html"
	ipSiteDoc := urlinfo.GetUrlDoc(ipUrl)
	return ipSiteDoc.Find("div#tab0_address").Text()
}

func GetRedisIpCount(ip string) (int64, error) {
	rdb := config.GetRedisCon()
	res, err := rdb.HGet("ip", ip).Int64()
	util.HandleError(err, "redis don't have ip "+ip)
	return res, err
}

func SetRedisIpCount(ip string, val int64) {
	rdb := config.GetRedisCon()
	rdb.HSet("ip", ip, val)
}

func UpdateIpInfoToMysql(ip string, count int64) {
	v2rayAccessLogIpCount := BuildUpdateIpInfo(ip, count)
	err := v2rayAccessLogIpCount.UpdateV2rayAccessLogIpCount(v2rayAccessLogIpCount)
	util.HandleError(err, "SaveIpInfoToMysql failed")
}

func SaveIpInfoToMysql(ip string, ipInfo string, count int64) {
	v2rayAccessLogIpCount := BuildSaveIpInfo(ip, ipInfo, count)
	err := v2rayAccessLogIpCount.SaveV2rayAccessLogIpCount(v2rayAccessLogIpCount)
	util.HandleError(err, "SaveIpInfoToMysql failed")
}

func BuildUpdateIpInfo(ip string, count int64) model.V2rayAccessLogIpCount {
	var v2rayAccessLogIpCount model.V2rayAccessLogIpCount
	v2rayAccessLogIpCount.Ip = ip
	v2rayAccessLogIpCount.Count = count
	v2rayAccessLogIpCount.UpdateTime = time.Now().Truncate(time.Second)
	return v2rayAccessLogIpCount
}

func BuildSaveIpInfo(ip string, ipInfo string, count int64) model.V2rayAccessLogIpCount {
	var v2rayAccessLogIpCount model.V2rayAccessLogIpCount
	v2rayAccessLogIpCount.Ip = ip
	v2rayAccessLogIpCount.IpInfo = ipInfo
	v2rayAccessLogIpCount.Count = count
	v2rayAccessLogIpCount.UpdateTime = time.Now().Truncate(time.Second)
	return v2rayAccessLogIpCount
}
