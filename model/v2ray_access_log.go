package model

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

func (V2rayAccessLog) TableName() string {
	return "v2ray_access_log"
}

func (v *V2rayAccessLog) SaveV2rayAccessLog(data V2rayAccessLog) error {
	return GetHhToolCon().Table(v.TableName()).Save(&data).Error
}

func (v *V2rayAccessLog) GetV2rayAccessLogById(minId int64, maxId int64) ([]V2rayAccessLog, error) {
	var v2rayAccessLogs []V2rayAccessLog
	err := GetHhToolCon().Table(v.TableName()).Where("id >= ? and id <= ?", minId, maxId).Find(&v2rayAccessLogs).Error
	return v2rayAccessLogs, err
}
