package bit_domain

import (
	"fmt"
	"hh_tool/model"
	"hh_tool/tool"
	"hh_tool/util"
	"time"

	"github.com/spf13/viper"
)

type BitDomainProcessor struct {
}

func NewBitDomainProcessor() (*BitDomainProcessor, error) {
	return &BitDomainProcessor{}, nil
}

func GetDomainInfo() {
	ipSite := viper.GetString("site.bit_domain.url")
	for {
		bitDomains, _ := new(model.BitDomain).GetNormalBitDomain(util.BitDomain_UnKnown)
		for _, bitDomain := range bitDomains {
			res := tool.GetBitDomainPost(ipSite, bitDomain.Domain+".bit")
			fmt.Println(bitDomain, res)
			if res.Err_No > 0 {
				continue
			}
			bitDomain.Status = res.Data.Status
			bitDomain.AccountPrice = res.Data.Account_Price
			bitDomain.BaseAmount = res.Data.Base_Amount
			bitDomain.UpdateTime = time.Now().Truncate(time.Second)
			new(model.BitDomain).UpdateBitDomainInfo(bitDomain)
		}
		if len(bitDomains) == 0 {
			break
		}
	}
	fmt.Println("GetDomainInfo finished")
}

func (p BitDomainProcessor) Monitor() {
	go MonitorSpecialBitDomainWithTask()
}

func MonitorSpecialBitDomainWithTask() {
	// 监控特别的域名
	bitDomainSpecicalSpec := viper.GetString("site.bit_domain_special.cron")
	cmd := func() {
		new(BitDomainProcessor).MonitorSpecialBitDomain()
	}
	util.SetCronTask(bitDomainSpecicalSpec, cmd)
}

func (p BitDomainProcessor) MonitorSpecialBitDomain() {
	ipSite := viper.GetString("site.bit_domain.url")
	bitDomains, _ := new(model.BitDomainSpecial).GetNotAvaliableBitDomainSpecial(util.BitDomain_Available_0, util.BitDomain_Available_1)
	for _, bitDomain := range bitDomains {
		bitUrl := bitDomain.Domain + ".bit"
		res := tool.GetBitDomainPost(ipSite, bitUrl)
		if res.Err_No > 0 {
			continue
		}
		if bitDomain.Status == res.Data.Status {
			bitDomain.MonitorCount += 1
			bitDomain.MonitorUpdateTime = time.Now().Truncate(time.Second)
		} else {
			bitDomain.Status = res.Data.Status
			bitDomain.AccountPrice = res.Data.Account_Price
			bitDomain.BaseAmount = res.Data.Base_Amount
			bitDomain.UpdateTime = time.Now().Truncate(time.Second)
			bitStatus := util.BitDomainMap[res.Data.Status]
			sendText := "[.bit域名监测] " + bitUrl + "的状态是" + bitStatus
			TelegramSendText(sendText)
		}
		new(model.BitDomainSpecial).UpdateBitDomainSpecialInfo(bitDomain)
	}
}
