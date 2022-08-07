package model

import (
	"hh_tool/config"
	"time"
)

type BitDomainSpecial struct {
	Id                int `gorm:"primary_key"`
	Domain            string
	Status            int
	AccountPrice      string
	BaseAmount        string
	UpdateTime        time.Time
	MonitorCount      int64
	MonitorUpdateTime time.Time
}

func (BitDomainSpecial) TableName() string {
	return "bit_domain_special"
}

func (v *BitDomainSpecial) GetNotAvaliableBitDomainSpecial(status int, status1 int) ([]BitDomainSpecial, error) {
	var BitDomainSpecial []BitDomainSpecial
	err := config.GetHhToolCon().Table(v.TableName()).Where("status not in(?,?) ", status, status1).Find(&BitDomainSpecial).Error
	return BitDomainSpecial, err
}

func (v *BitDomainSpecial) UpdateBitDomainSpecialInfo(data BitDomainSpecial) error {
	return config.GetHhToolCon().Table(v.TableName()).Select("status", "account_price", "base_amount", "update_time", "monitor_count", "monitor_update_time").Updates(data).Error
}
