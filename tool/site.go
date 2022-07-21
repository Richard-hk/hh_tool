package tool

import (
	"bytes"
	"encoding/json"
	"hh_tool/util"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetSiteDoc(url string) *goquery.Document {
	resp, err := http.Get(url)
	util.HandleError(err, "GetSiteStr failed")
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	util.HandleError(err, "GetSiteDoc failed")
	return doc
}

type BitDomain struct {
	Account          string `json:"account"`
	Address          string `json:"address"`
	Account_Char_Str []AccountCharStr
}
type AccountCharStr struct {
	Char_Set_Name int    `json:"char_set_name"`
	Char          string `json:"char"`
}

type RespBitDomain struct {
	Err_No  int
	Err_Msg string
	Data    RespBitDomainData
}

type RespBitDomainData struct {
	Status        int
	Account       string
	Account_Price string
	Base_Amount   string
}

func GetBitDomainPost(url string, account string) (res RespBitDomain) {
	accountCharStr := []AccountCharStr{}
	for _, val := range account {
		charSetName := 1
		if !IsSingleDigit(string(val)) {
			charSetName = 2
		}
		tmpAccountChartStr := AccountCharStr{Char_Set_Name: charSetName, Char: string(val)}
		accountCharStr = append(accountCharStr, tmpAccountChartStr)
	}
	bitDomain := &BitDomain{Account: account, Address: "", Account_Char_Str: accountCharStr}
	json_data, _ := json.Marshal(bitDomain)
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(json_data))
	util.HandleError(err, "GetSiteStr failed")
	json.NewDecoder(resp.Body).Decode(&res)
	return res
}
