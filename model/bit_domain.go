package model

import "time"

type BitDomain struct {
	Id           int `gorm:"primary_key"`
	Domain       string
	Status       int
	AccountPrice string
	BaseAmount   string
	UpdateTime   time.Time
}

func (BitDomain) TableName() string {
	return "bit_domain"
}

func (v *BitDomain) GetNormalBitDomain(status int) ([]BitDomain, error) {
	var BitDomain []BitDomain
	err := GetHhToolCon().Table(v.TableName()).Where("status = ?", status).Find(&BitDomain).Error
	return BitDomain, err
}

func (v *BitDomain) UpdateBitDomainInfo(data BitDomain) error {
	return GetHhToolCon().Table(v.TableName()).Select("status", "account_price", "base_amount", "update_time").Updates(data).Error
}
