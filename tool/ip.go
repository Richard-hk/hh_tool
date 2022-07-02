package tool

import (
	"hh_tool/database"
	"hh_tool/model"
	"hh_tool/util"
	"time"

	"github.com/spf13/viper"
)

func SaveIpInfo(v2rayIpCountMap map[string]int) {
	for ip, count := range v2rayIpCountMap {
		redisInfoCount, err := GetRedisIpCount(ip)
		if err == nil {
			count := redisInfoCount + int64(count)
			SetRedisIpCount(ip, count)
			UpdateIpInfoToMysql(ip, count)
		} else {
			ipInfo := GetIpInfo(ip)
			SetRedisIpCount(ip, int64(count))
			SaveIpInfoToMysql(ip, ipInfo, int64(count))
		}
	}
}

func GetIpInfo(ip string) string {
	ipSite := viper.GetString("site.ip.url")
	ipUrl := ipSite + "/" + ip + ".html"
	ipSiteDoc := GetSiteDoc(ipUrl)
	return ipSiteDoc.Find("div#tab0_address").Text()
}

func GetRedisIpCount(ip string) (int64, error) {
	rdb := database.GetRedisCon()
	res, err := rdb.HGet("ip", ip).Int64()
	util.HandleError(err, "redis don't have ip "+ip)
	return res, err
}

func SetRedisIpCount(ip string, val int64) {
	rdb := database.GetRedisCon()
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
