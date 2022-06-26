package model

import "time"

type V2rayAccessLogIpCount struct {
	ID         int `gorm:"primary_key"`
	Ip         string
	IpInfo     string
	Count      int64
	UpdateTime time.Time
}

func (*V2rayAccessLogIpCount) TableName() string {
	return "v2ray_access_log_ip_count"
}

func (v *V2rayAccessLogIpCount) SaveV2rayAccessLogIpCount(data V2rayAccessLogIpCount) error {
	return GetHhToolCon().Table(v.TableName()).Save(&data).Error
}

func (v *V2rayAccessLogIpCount) UpdateV2rayAccessLogIpCount(data V2rayAccessLogIpCount) error {
	return GetHhToolCon().Table(v.TableName()).Where("ip", data.Ip).Updates(data).Error
}
