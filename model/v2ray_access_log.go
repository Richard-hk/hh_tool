package model

type V2rayAccessLog struct {
	ID            int `gorm:"primary_key"`
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

func (V2rayAccessLog) TableName() string {
	return "v2ray_access_log"
}

func (v *V2rayAccessLog) SaveV2rayAccessLog(data V2rayAccessLog) error {
	return GetHhToolCon().Table(v.TableName()).Save(&data).Error
}
