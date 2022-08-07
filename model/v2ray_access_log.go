package model

import "hh_tool/config"

type V2rayAccessLog struct {
	Id            int `gorm:"primary_key"`
	Dt            string
	Time          string
	Ip            string
	Port          string
	Status        string
	RemoteAdr     string
	RemoteAdrPort string
	Type          string
	Reason        string
}

type V2rayAccessLogIpCountNotice struct {
	Ip    string
	Count int64
}

func (V2rayAccessLog) TableName() string {
	return "v2ray_access_log"
}

func (v *V2rayAccessLog) SaveV2rayAccessLog(data V2rayAccessLog) error {
	return config.GetHhToolCon().Table(v.TableName()).Save(&data).Error
}

func (v *V2rayAccessLog) GetV2rayAccessLogById(minId int64, maxId int64) ([]V2rayAccessLog, error) {
	var v2rayAccessLogs []V2rayAccessLog
	err := config.GetHhToolCon().Table(v.TableName()).Where("id >= ? and id <= ?", minId, maxId).Find(&v2rayAccessLogs).Error
	return v2rayAccessLogs, err
}

func (v *V2rayAccessLog) GetV2rayAccessLogIpCountByDt(dt int64, status string) ([]V2rayAccessLogIpCountNotice, error) {
	var v2rayAccessLogIpCountNotice []V2rayAccessLogIpCountNotice
	err := config.GetHhToolCon().Table(v.TableName()).Select("ip, count(ip) count").Where("dt = ? and `status` = ?", dt, status).Group("ip").Find(&v2rayAccessLogIpCountNotice).Error
	return v2rayAccessLogIpCountNotice, err
}
