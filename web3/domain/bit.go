package domain

import (
	"fmt"
	"hh_tool/model"
	"hh_tool/tool"
	"hh_tool/util"
	"time"

	"github.com/spf13/viper"
)

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
