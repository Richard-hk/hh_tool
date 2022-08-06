package model

import "time"

type BitDomainSpecial struct {
	Id           int `gorm:"primary_key"`
	Domain       string
	Status       int
	AccountPrice string
	BaseAmount   string
	UpdateTime   time.Time
}

func (BitDomainSpecial) TableName() string {
	return "bit_domain_special"
}

func (v *BitDomainSpecial) GetNotAvaliableBitDomainSpecial(status int, status1 int) ([]BitDomainSpecial, error) {
	var BitDomainSpecial []BitDomainSpecial
	err := GetHhToolCon().Table(v.TableName()).Where("status not in(?,?) ", status, status1).Find(&BitDomainSpecial).Error
	return BitDomainSpecial, err
}

func (v *BitDomainSpecial) UpdateBitDomainSpecialInfo(data BitDomainSpecial) error {
	return GetHhToolCon().Table(v.TableName()).Select("status", "account_price", "base_amount", "update_time").Updates(data).Error
}
